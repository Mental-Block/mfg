package services

import (
	"context"
	"fmt"

	"github.com/server/internal"
	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)

type ResourceService struct {
	resourceStore ports.ResourceStore
}

func NewResourceService(resource ports.ResourceStore) *ResourceService {
	return &ResourceService{
		resourceStore: resource,
	}
}

func (s *ResourceService) GetResource(ctx context.Context, id int) (*domain.Resource, error) {
	validId := domain.NewId(id)

	resource, err := s.resourceStore.Select(ctx, validId)

	if (err != nil) {
		if err.Error() == domain.ErrResourceNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error())
		}
	
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrResourceStore.Error())
	}

	return resource, nil
}
	
func (s *ResourceService) GetResources(ctx context.Context) ([]domain.Resource, error)  {
	resources, err := s.resourceStore.SelectResources(ctx)

	if (err != nil) {
		if err.Error() == domain.ErrResourceNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrResourceStore.Error())
	}

	return resources, nil
}

func (s *ResourceService) New(ctx context.Context, name string) (*domain.Resource, error) {

	validResource, err := domain.NewResource(name);

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid resource supplied: %s", err.Error()))
	}

	resource, err := s.resourceStore.Insert(ctx, validResource)

	if (err != nil) {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrResourceStore.Error())
	}

	return resource, nil
}

func (s *ResourceService) Remove(ctx context.Context, id int) (*domain.Id, error) {
	
	validId := domain.NewId(id)

	resourceId, err := s.resourceStore.Delete(ctx, validId)

	if (err != nil) {
		if err.Error() == domain.ErrRoleNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrResourceStore.Error())
	}

	return resourceId, nil
}

func (s *ResourceService) Update(ctx context.Context, id int, name string) (*domain.Resource, error) {
	validId := domain.NewId(id)

	validResource, err := domain.NewResource(name);

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid resource supplied: %s", err.Error()))
	}

	resource, err := s.resourceStore.Update(ctx, validId, validResource)
		
	if (err != nil) {
		if err.Error() == domain.ErrResourceNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrResourceStore.Error())
	}

	return resource, nil
}

