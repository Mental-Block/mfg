package spicedb

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync/atomic"

	authzedV1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/server/internal/core/relation/domain"
	"github.com/server/pkg/logger"
)

/*
 High level overview of ISpiceRelationStore should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type ISpiceRelationStore interface {
	Check(ctx context.Context, rel domain.Relation) (bool, error)
	BatchCheck(ctx context.Context, relations []domain.Relation) ([]domain.CheckPair, error)
	Delete(ctx context.Context, rel domain.Relation) error
	Add(ctx context.Context, rel domain.Relation) error
	LookupSubjects(ctx context.Context, rel domain.Relation) ([]string, error)
	LookupResources(ctx context.Context, rel domain.Relation) ([]string, error)
	ListRelations(ctx context.Context, rel domain.Relation) ([]domain.Relation, error)
}

type RelationStore struct {
	spiceDB *SpiceDB

	// Consistency ensures Authz server consistency guarantees for various operations
	// Possible values are:
	// - "full": Guarantees that the data is always fresh
	// - "best_effort": Guarantees that the data is the best effort fresh
	// - "minimize_latency": Tries to prioritise minimal latency
	consistency ConsistencyLevel

	// tracing enables debug traces for check calls
	tracing bool

	// lastToken is the last zookie returned by the server, this is cached at instance level and
	// maybe not be consistent across multiple instances but that is fine in most cases as
	// the token is only used in lookup or list calls, for permission checks we always use the
	// consistency level. Storing it in a shared db/cache will make it consistent across instances.
	// We can also store multiple tokens in the cache based on what kind of resource we are dealing with
	// but that adds complexity.
	lastToken atomic.Pointer[authzedV1.ZedToken]

	logger logger.Logger
}

func NewRelationStore(
	spiceDB *SpiceDB, 
	consistency string, 
	tracing bool,
	logr logger.Logger,
	) (*RelationStore) {
	return &RelationStore{
		spiceDB:     spiceDB,
		consistency: ConsistencyLevel(consistency),
		tracing:     tracing,
		logger: 	 logr,
	}
}

func (r *RelationStore) Add(ctx context.Context, rel domain.Relation) error {
	relationship := &authzedV1.Relationship{
		Resource: &authzedV1.ObjectReference{
			ObjectId:   rel.Object.Id,
			ObjectType: rel.Object.Namespace,
		},
		Relation: rel.RelationName,
		Subject: &authzedV1.SubjectReference{
			Object: &authzedV1.ObjectReference{
				ObjectId:   rel.Subject.Id,
				ObjectType: rel.Subject.Namespace,
			},
			OptionalRelation: rel.Subject.SubRelationName,
		},
	}

	request := &authzedV1.WriteRelationshipsRequest{
		Updates: []*authzedV1.RelationshipUpdate{
			{
				Operation:    authzedV1.RelationshipUpdate_OPERATION_TOUCH,
				Relationship: relationship,
			},
		},
	}

	resp, err := r.spiceDB.Client.WriteRelationships(ctx, request)
	
	if err != nil {
		return err
	}

	r.lastToken.Store(resp.GetWrittenAt())
	return nil
}

func (r *RelationStore) Check(ctx context.Context, rel domain.Relation) (bool, error) {
	request := &authzedV1.CheckPermissionRequest{
		Consistency: r.getConsistencyForCheck(),
		Resource: &authzedV1.ObjectReference{
			ObjectId:   rel.Object.Id,
			ObjectType: rel.Object.Namespace,
		},
		
		Subject: &authzedV1.SubjectReference{
			Object: &authzedV1.ObjectReference{
				ObjectId:   rel.Subject.Id,
				ObjectType: rel.Subject.Namespace,
			},
			OptionalRelation: rel.Subject.SubRelationName,
		},
		Permission:  rel.RelationName,
		WithTracing: r.tracing,
	}


	response, err := r.spiceDB.Client.CheckPermission(ctx, request)
	if err != nil {
		return false, err
	}

	if response.GetDebugTrace() != nil {
		str, _ := json.Marshal(response.GetDebugTrace())
		
		logger.UnWrapGRPCCtx(ctx).Info("CheckPermission", logger.ZapString("trace", string(str)))
	}

	r.lastToken.Store(response.GetCheckedAt())
	return response.GetPermissionship() == authzedV1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, nil
}

func (r *RelationStore) Delete(ctx context.Context, rel domain.Relation) error {
	if rel.Object.Namespace == "" {
		return errors.New("object namespace is required to delete a relation")
	}
	request := &authzedV1.DeleteRelationshipsRequest{
		RelationshipFilter: &authzedV1.RelationshipFilter{
			ResourceType:       rel.Object.Namespace,
			OptionalResourceId: rel.Object.Id,
			OptionalRelation:   rel.RelationName,
			OptionalSubjectFilter: &authzedV1.SubjectFilter{
				SubjectType:       rel.Subject.Namespace,
				OptionalSubjectId: rel.Subject.Id,
				OptionalRelation: &authzedV1.SubjectFilter_RelationFilter{
					Relation: rel.Subject.SubRelationName,
				},
			},
		},
	}

	resp, err := r.spiceDB.Client.DeleteRelationships(ctx, request)

	if err != nil {
		return err
	}

	r.lastToken.Store(resp.GetDeletedAt())

	return nil
}

func (r *RelationStore) LookupSubjects(ctx context.Context, rel domain.Relation) ([]string, error) {
	resp, err := r.spiceDB.Client.LookupSubjects(ctx, &authzedV1.LookupSubjectsRequest{
		Consistency: r.getConsistency(),
		Resource: &authzedV1.ObjectReference{
			ObjectType: rel.Object.Namespace,
			ObjectId:   rel.Object.Id,
		},
		Permission:        rel.RelationName,
		SubjectObjectType: rel.Subject.Namespace,
	})

	if err != nil {
		return nil, err
	}

	var subjects []string

	for {
		item, err := resp.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		subjects = append(subjects, item.GetSubject().GetSubjectObjectId())
	}

	return subjects, nil
}

func (r *RelationStore) LookupResources(ctx context.Context, rel domain.Relation) ([]string, error) {
	
	resp, err := r.spiceDB.Client.LookupResources(ctx, &authzedV1.LookupResourcesRequest{
		Consistency:        r.getConsistency(),
		ResourceObjectType: rel.Object.Namespace,
		Permission:         rel.RelationName,
		Subject: &authzedV1.SubjectReference{
			Object: &authzedV1.ObjectReference{
				ObjectType: rel.Subject.Namespace,
				ObjectId:   rel.Subject.Id,
			},
			OptionalRelation: rel.Subject.SubRelationName,
		},
	})

	if err != nil {
		return nil, err
	}

	var subjects []string

	for {
		item, err := resp.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		subjects = append(subjects, item.GetResourceObjectId())
	}

	return subjects, nil
}

// ListRelations shouldn't be used in high TPS flows as consistency requirements are set high
func (r *RelationStore) ListRelations(ctx context.Context, rel domain.Relation) ([]domain.Relation, error) {
	
	resp, err := r.spiceDB.Client.ReadRelationships(ctx, &authzedV1.ReadRelationshipsRequest{
		Consistency: r.getConsistency(),
		RelationshipFilter: &authzedV1.RelationshipFilter{
			ResourceType:       rel.Object.Namespace,
			OptionalResourceId: rel.Object.Id,
			OptionalRelation:   rel.RelationName,
			OptionalSubjectFilter: &authzedV1.SubjectFilter{
				SubjectType:       rel.Subject.Namespace,
				OptionalSubjectId: rel.Subject.Id,
				OptionalRelation:  nil,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	var rels []domain.Relation

	for {
		item, err := resp.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		pbRel := item.GetRelationship()

		rels = append(rels, domain.Relation{
			Object: domain.Object{
				Id: pbRel.GetResource().GetObjectId(),
				Namespace: pbRel.GetResource().GetObjectType(),
			},
			Subject: domain.Subject{
				Id: pbRel.GetSubject().GetObject().GetObjectId(),
				Namespace: pbRel.GetSubject().GetObject().GetObjectType(),
				SubRelationName: pbRel.GetRelation(),
			},
		})
	}

	return rels, nil
}

func (r *RelationStore) BatchCheck(ctx context.Context, relations []domain.Relation) ([]domain.CheckPair, error) {
	
	result := make([]domain.CheckPair, len(relations))

	items := make([]*authzedV1.CheckBulkPermissionsRequestItem, 0, len(relations))

	for _, rel := range relations {
		items = append(items, &authzedV1.CheckBulkPermissionsRequestItem{
			Resource: &authzedV1.ObjectReference{
				ObjectId:   rel.Object.Id,
				ObjectType: rel.Object.Namespace,
			},
			Subject: &authzedV1.SubjectReference{
				Object: &authzedV1.ObjectReference{
					ObjectId:   rel.Subject.Id,
					ObjectType: rel.Subject.Namespace,
				},
				OptionalRelation: rel.Subject.SubRelationName,
			},
			Permission: rel.RelationName,
		})
	}
	
	request := &authzedV1.CheckBulkPermissionsRequest{
		Consistency: r.getConsistencyForCheck(),
		Items:       items,
	}

	response, err := r.spiceDB.Client.CheckBulkPermissions(ctx, request)

	if err != nil {
		return result, err
	}

	var respErr error = nil
	for itemIdx, item := range response.GetPairs() {
		result[itemIdx] = domain.CheckPair{
			Relation: domain.Relation{
				Object: domain.Object{
					Id: item.GetRequest().GetResource().GetObjectId(),
					Namespace: item.GetRequest().GetResource().GetObjectType(),
				},
				Subject: domain.Subject{
					Id: item.GetRequest().GetSubject().GetObject().GetObjectId(),
					Namespace: item.GetRequest().GetSubject().GetObject().GetObjectType(),
					SubRelationName:  item.GetRequest().GetSubject().GetOptionalRelation(),
				},
				RelationName: item.GetRequest().GetPermission(),
			},
			Status: false,
		}

		if item.GetError() != nil {
			respErr = errors.Join(respErr, errors.New(item.GetRequest().GetPermission()+": "+item.GetError().GetMessage()))
			continue
		}

		if item.GetItem() != nil {
			result[itemIdx].Status = item.GetItem().GetPermissionship() == authzedV1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION
		}
	}

	r.lastToken.Store(response.GetCheckedAt())
	return result, respErr
}

func (r *RelationStore) getConsistency() *authzedV1.Consistency {
	switch r.consistency {
		case ConsistencyLevelMinimizeLatency:
			return &authzedV1.Consistency{Requirement: &authzedV1.Consistency_MinimizeLatency{MinimizeLatency: true}}
		case ConsistencyLevelFull:
			return &authzedV1.Consistency{Requirement: &authzedV1.Consistency_FullyConsistent{FullyConsistent: true}}
	}

	lastToken := r.lastToken.Load()

	if lastToken == nil {
		return &authzedV1.Consistency{Requirement: &authzedV1.Consistency_FullyConsistent{FullyConsistent: true}}
	}

	return &authzedV1.Consistency{
		Requirement: &authzedV1.Consistency_AtLeastAsFresh{
			AtLeastAsFresh: lastToken,
		},
	}
}

func (r *RelationStore) getConsistencyForCheck() *authzedV1.Consistency {

	if r.consistency == ConsistencyLevelMinimizeLatency {
		return &authzedV1.Consistency{Requirement: &authzedV1.Consistency_MinimizeLatency{MinimizeLatency: true}}
	}

	return &authzedV1.Consistency{Requirement: &authzedV1.Consistency_FullyConsistent{FullyConsistent: true}}
}
