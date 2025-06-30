package domain

import (
	namespace "github.com/server/internal/core/namespace/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type SanitizedPermission struct {
	Id          utils.UUID
	Name        PermissionName
	Slug        PermissionSlug
	NamespaceId namespace.NameSpaceName
	Metadata    metadata.Metadata		
}

func (r SanitizedPermission) Transform()(*Permission){
	return &Permission{
		Id: r.Id.String(),
		Name: r.Name.String(),
		Slug: r.Slug.String(),
		NamespaceId: r.NamespaceId.String(),
		Metadata: r.Metadata,
	}
}

type Permission struct {
	Id          string
	Name        string
	Slug        string
	NamespaceId string
	Metadata    metadata.Metadata		
}

func (r Permission) Transform()(*SanitizedPermission, error) {

	if (r.Id == "") {
		r.Id = utils.NewUUID().String()
	}

	UUID, err := utils.ConvertStringToUUID(r.Id) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	perm := PermissionName(r.Name)

	if err = perm.IsValid(); err != nil {
		return nil, err
	}

	ns := namespace.NameSpaceName(r.NamespaceId)

	if err := ns.IsValid(); err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, err.Error()) 
	}

	if (r.Slug == "") {
		r.Slug = perm.BuildSlug(ns).String()
	}
 
	slug := PermissionSlug(r.Slug)

	if err := slug.IsValid(); err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, err.Error()) 
	}

	return &SanitizedPermission{
		Id: UUID,
		Name: perm,
		NamespaceId: ns,
		Slug: slug,
		Metadata: r.Metadata,
	}, nil
}


// add namespace resource paramter to namespace if permission has one of the following dilimeters: NamespaceDilimeter, PermissionDilimeter or PermissionKeyDilimeter.
// example: app/resource, my#permission -> app/resource:my#permission
// func (s NameSpaceName) AddResourceParamIfRequired(permission string) string {
// 	if  strings.Contains(s.String(), utils.RelationDilimeter) || 
// 		strings.Contains(s.String(), utils.KeyDilimeter) || 
// 		strings.Contains(s.String(), utils.NamespaceDilimeter) {
// 		return s.String()
// 	}

// 	return s.JoinParam(s.String())
// }