package services

import (
	"context"
	"fmt"

	"github.com/server/internal"
	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)

type RoleService struct {
	roleStore ports.RoleStore
}

func NewRoleService(role ports.RoleStore) *RoleService {
	return &RoleService{
		roleStore: role,
	}
}

func (s *RoleService) GetRole(ctx context.Context, id int) (*domain.Role, error) {
	validId := domain.NewId(id)

	role, err := s.roleStore.Select(ctx, validId)

	if (err != nil) {
		if err.Error() == domain.ErrRoleNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error())
		}
	
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrRoleStore.Error())
	}

	return role, nil
}
	
func (s *RoleService) GetRoles(ctx context.Context) ([]domain.Role, error)  {
	roles, err := s.roleStore.SelectRoles(ctx)

	if (err != nil) {
		if err.Error() == domain.ErrRoleNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrRoleStore.Error())
	}

	return roles, nil
}

func (s *RoleService) New(ctx context.Context, name string) (*domain.Role, error) {

	validRole, err := domain.NewRole(name);

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid role supplied: %s", err.Error()))
	}

	role, err := s.roleStore.Insert(ctx, validRole)

	if (err != nil) {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrRoleStore.Error())
	}

	return role, nil
}

func (s *RoleService) Remove(ctx context.Context, id int) (*domain.Id, error) {
	
	validId := domain.NewId(id)

	roleId, err := s.roleStore.Delete(ctx, validId)

	if (err != nil) {
		if err.Error() == domain.ErrRoleNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrRoleStore.Error())
	}

	return roleId, nil
}

func (s *RoleService) Update(ctx context.Context, id int, name string) (*domain.Role, error) {
	validId := domain.NewId(id)

	validRole, err := domain.NewRole(name);

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid role supplied: %s", err.Error()))
	}

	roleId, err := s.roleStore.Update(ctx, validId, validRole)

	if (err != nil) {
		if err.Error() == domain.ErrRoleNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrRoleStore.Error())
	}

	return roleId, nil
}

