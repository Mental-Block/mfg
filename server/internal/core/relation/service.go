package relation

import (
	"context"

	"github.com/server/internal/core/relation/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of RelationService should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IRelationService interface {
	BatchCheckPermission(ctx context.Context, relations []domain.Relation) ([]domain.CheckPair, error)
	CheckPermission(ctx context.Context, rel domain.Relation) (bool, error)
	Get(ctx context.Context, id string) (*domain.Relation, error)
	List(ctx context.Context, flt domain.Filter) ([]domain.Relation, error)
	ListByFields(ctx context.Context, input domain.Relation) ([]domain.Relation, error)
	ListRelations(ctx context.Context, rel domain.Relation) ([]domain.Relation, error)
	LookupSubjects(ctx context.Context, rel domain.Relation) ([]string, error)
	LookupResources(ctx context.Context, rel domain.Relation) ([]string, error)
	New(ctx context.Context, input domain.Relation) (*domain.Relation, error)
 	Remove(ctx context.Context, rel domain.Relation) error
}

type RelationService struct {
	relationStore  IRelationStore
	spiceRelationStore ISpiceRelationStore
}

func NewRelationService(
	relationStore IRelationStore, 
	spiceRelationStore ISpiceRelationStore,
) *RelationService {
	return &RelationService{
		relationStore:      	relationStore,
		spiceRelationStore: 	spiceRelationStore,
	}
}

func (s RelationService) Get(ctx context.Context, id string) (*domain.Relation, error) {

	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	model, err := s.relationStore.Select(ctx, UUID)

	if err != nil {	
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrRelationService)
	}

	relation := model.Transform().Transform()

	return relation, nil
}

func (s RelationService) New(ctx context.Context, input domain.Relation) (*domain.Relation, error) {

	input.Id =  utils.NewUUID().String()

	sanitizedRelation, err := input.Transform()
	
	if err != nil {
		return nil, err
	}

	relation := sanitizedRelation.Transform()

	err = s.spiceRelationStore.Add(ctx, *relation)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrCreatingRelationInAuthzEngine.Error())
	}

	_, err = s.relationStore.Upsert(ctx, sanitizedRelation)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrCreatingRelationInStore.Error())
	}

	return relation, nil
}

func (s RelationService) List(ctx context.Context, flt domain.Filter) ([]domain.Relation, error) {

	models, err := s.relationStore.Selects(ctx)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrRelationService)
	}

	relations := make([]domain.Relation, len(models))

	for i, _ := range relations {
		relations[i] = *models[i].Transform().Transform()
	}

	return relations, nil
}

func (s RelationService) ListByFields(ctx context.Context, input domain.Relation) ([]domain.Relation, error) {

	sanitizedRelation, err := input.Transform()

	if err != nil {
		return nil, err
	}

	models, err := s.relationStore.SelectsByFields(ctx, sanitizedRelation)

	relations := make([]domain.Relation, len(models))

	for i, _ := range relations {
		relations[i] = *models[i].Transform().Transform()
	}

	return relations, nil
}

func (s RelationService) Remove(ctx context.Context, rel domain.Relation) error {

	fields, err := s.ListByFields(ctx, rel)

	if err != nil {
		return err
	}

	for _, field := range fields {
		if err = s.spiceRelationStore.Delete(ctx, field); err != nil {
			return err
		}
		
		// we can safely cast this to utils.UUID without checking as it is returned by s.ListByFields call that does the validation
		if _, err = s.relationStore.Delete(ctx, utils.UUID(field.Id)); err != nil {
			return err
		}
	}

	return nil
}

func (s RelationService) CheckPermission(ctx context.Context, input domain.Relation) (bool, error) {

	sanitizedRelation, err := input.Transform()

	if (err != nil) {
		return false, nil
	}

	rel := *sanitizedRelation.Transform()

	return s.spiceRelationStore.Check(ctx, rel)
}

func (s RelationService) BatchCheckPermission(ctx context.Context, input []domain.Relation) ([]domain.CheckPair, error) {

	relations := make([]domain.Relation, len(input))

	for i, _ := range relations {
		 santizedRel, err := relations[i].Transform()
		 
		 if (err != nil) {
			return nil, err
		 }
		 
		 relations[i] = *santizedRel.Transform()
	}

	return s.spiceRelationStore.BatchCheck(ctx, relations)
}

// LookupSubjects returns all the subjects of a given type that have access whether
// via a computed permission or relation membership.
func (s RelationService) LookupSubjects(ctx context.Context, input domain.Relation) ([]string, error) {
	
	santizedRel, err := input.Transform()
		 
	if (err != nil) {
		return nil, err
	}
		 
	rel := *santizedRel.Transform()

	return s.spiceRelationStore.LookupSubjects(ctx, rel)
}

// LookupResources returns all the resources of a given type that a subject can access whether
// via a computed permission or relation membership.
func (s RelationService) LookupResources(ctx context.Context, input domain.Relation) ([]string, error) {
	
	santizedRel, err := input.Transform()
		 
	if (err != nil) {
		return nil, err
	}
		 
	rel := *santizedRel.Transform()

	return s.spiceRelationStore.LookupResources(ctx, rel)
}

// ListRelations lists a set of the relationships matching filter
func (s RelationService) ListRelations(ctx context.Context, input domain.Relation) ([]domain.Relation, error) {

	santizedRel, err := input.Transform()
		 
	if (err != nil) {
		return nil, err
	}
		 
	rel := *santizedRel.Transform()

	return s.spiceRelationStore.ListRelations(ctx, rel)
}