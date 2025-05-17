package services

import (
	"context"
	"fmt"

	"github.com/server/internal"
	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)

type PermissionService struct {
	permissionStore ports.PermissionStore
}

func NewPermissionService(permission ports.PermissionStore) *PermissionService {
	return &PermissionService{
		permissionStore: permission,
	}
}

func (s *PermissionService) GetPermission(ctx context.Context, id int) (*domain.Permission, error) {
	validId := domain.NewId(id)

	permission, err := s.permissionStore.Select(ctx, validId)

	if (err != nil) {
		if err.Error() == domain.ErrPermissionNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error())
		}
	
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrPermissionStore.Error())
	}

	return permission, nil
}
	
func (s *PermissionService) GetPermissions(ctx context.Context) ([]domain.Permission, error)  {
	permissions, err := s.permissionStore.SelectPermissions(ctx)

	if (err != nil) {
		if err.Error() == domain.ErrPermissionNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrPermissionStore.Error())
	}

	return permissions, nil
}

func (s *PermissionService) New(ctx context.Context, name string) (*domain.Permission, error) {

	validPermission, err := domain.NewPermission(name);

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid permission supplied: %s", err.Error()))
	}

	permission, err := s.permissionStore.Insert(ctx, validPermission)

	if (err != nil) {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrPermissionStore.Error())
	}

	return permission, nil
}

func (s *PermissionService) Remove(ctx context.Context, id int) (*domain.Id, error) {
	
	validId := domain.NewId(id)

	permisionId, err := s.permissionStore.Delete(ctx, validId)

	if (err != nil) {
		if err.Error() == domain.ErrPermissionNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrPermissionStore.Error())
	}

	return permisionId, nil
}

func (s *PermissionService) Update(ctx context.Context, id int, name string) (*domain.Permission, error) {
	validId := domain.NewId(id)

	validPermission, err := domain.NewPermission(name);

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid permission supplied: %s", err.Error()))
	}

	permission, err := s.permissionStore.Update(ctx, validId, validPermission)

	if (err != nil) {
		if err.Error() == domain.ErrPermissionNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrPermissionStore.Error())
	}

	return permission, nil
}

