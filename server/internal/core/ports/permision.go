package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type PermissionStore interface {
	Delete(ctx context.Context, id domain.Id) (*domain.Id, error)
	Insert(ctx context.Context, name string) (*domain.Permission, error)
	Select(ctx context.Context, id domain.Id) (*domain.Permission, error)
	SelectPermissions(ctx context.Context) ([]domain.Permission, error) 
	Update(ctx context.Context, id domain.Id, name string) (*domain.Permission, error)
}

type PermissionService interface {
	GetPermission(ctx context.Context, id int) (*domain.Permission, error)
	GetPermissions(ctx context.Context) ([]domain.Permission, error)
	New(ctx context.Context, name string) (*domain.Permission, error)
	Remove(ctx context.Context, id int) (*domain.Id, error)
	Update(ctx context.Context, id int, name string) (*domain.Permission, error)
}