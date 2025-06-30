package permission

import (
	"context"

	permission_store "github.com/server/internal/adapters/store/postgres/permission"
	"github.com/server/internal/core/permission/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of IPermissionService should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IPermissionService interface { 
	//get by id or slug
	Get(ctx context.Context, id string) (*domain.Permission, error) 
	Upsert(ctx context.Context, perm domain.Permission) (*domain.Permission, error) 
	List(ctx context.Context) ([]domain.Permission, error)
	Update(ctx context.Context, perm domain.Permission) (*domain.Permission, error)
	Remove(ctx context.Context, id string) (string, error)
}

type PermissionService struct {
	permissionStore IPermissionStore
}

func NewService(store IPermissionStore) *PermissionService {
	return &PermissionService{
		permissionStore: store,
	}
}

func (s PermissionService) Get(ctx context.Context, id string) (*domain.Permission, error) {

	var models *permission_store.PermissionModel
	var err error

	if utils.IsValidUUID(id) {
		models, err = s.permissionStore.Select(ctx, utils.UUID(id))
	} else {
		models, err = s.permissionStore.SelectBySlug(ctx, domain.ToSlug(id))
	}

	if (err != nil) {
		return nil, err
	}

	sanitized, err := models.Transform()

	if err != nil {
		return nil, err
	}

	permission := sanitized.Transform()

	return permission, nil
}

func (s PermissionService) Upsert(ctx context.Context, perm domain.Permission) (*domain.Permission, error) {
	
	santized, err := perm.Transform()

	if (err != nil) {
		return nil, err
	}

	model, err := s.permissionStore.Upsert(ctx, *santized)

	if (err != nil) {
		return nil, err
	}

	sanitizedPerm, err := model.Transform()

	if (err != nil) {
		return nil, err
	}

	permission := sanitizedPerm.Transform()

	return permission, nil
}

func (s PermissionService) List(ctx context.Context) ([]domain.Permission, error) {

	models, err := s.permissionStore.Selects(ctx)

	if (err != nil) {
		return nil, err
	}

	permissions := make([]domain.Permission, len(models))

	for i := range permissions {
		santized, err := models[i].Transform()
		
		if (err != nil) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrPermissionService)
		}

		permissions[i] = *santized.Transform()
	}

	return permissions, nil
}

func (s PermissionService) Update(ctx context.Context, perm domain.Permission) (*domain.Permission, error) {
	santized, err := perm.Transform()

	if (err != nil) {
		return nil, err
	}

	model, err := s.permissionStore.Update(ctx, *santized)

	if (err != nil) {
		return nil, err
	}

	sanitizedPerm, err := model.Transform()

	if (err != nil) {
		return nil, err
	}

	permission := sanitizedPerm.Transform()

	return permission, nil
}

// Delete call over a service could be dangerous without removing all of its relations
// the method does not do it by default
func (s PermissionService) Remove(ctx context.Context, id string) (string, error) {
	
	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return "", utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	dataId, err := s.permissionStore.Delete(ctx, UUID)

	if err != nil {
		return "", utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrPermissionService)
	}

	return string(*dataId), nil
}
