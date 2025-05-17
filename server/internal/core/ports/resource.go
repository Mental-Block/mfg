package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type ResourceStore interface {
	Delete(ctx context.Context, id domain.Id) (*domain.Id, error)
	Insert(ctx context.Context, name string) (*domain.Resource, error)
	Select(ctx context.Context, id domain.Id) (*domain.Resource, error)
	SelectResources(ctx context.Context) ([]domain.Resource, error) 
	Update(ctx context.Context, id domain.Id, name string) (*domain.Resource, error)
}

type ResourceService interface {
	GetResource(ctx context.Context, id int) (*domain.Resource, error)
	GetResources(ctx context.Context) ([]domain.Resource, error)
	New(ctx context.Context, name string) (*domain.Resource, error)
	Remove(ctx context.Context, id int) (*domain.Id, error)
	Update(ctx context.Context, id int, name string) (*domain.Resource, error)
}