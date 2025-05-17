package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type RoleStore interface {
	Delete(ctx context.Context, id domain.Id) (*domain.Id, error)
	Insert(ctx context.Context, name string) (*domain.Role, error)
	Select(ctx context.Context, id domain.Id) (*domain.Role, error)
	Update(ctx context.Context, id domain.Id, name string) (*domain.Role, error)
	SelectRoles(ctx context.Context) ([]domain.Role, error) 
}

type RoleService interface {
	GetRole(ctx context.Context, id int) (*domain.Role, error)
	GetRoles(ctx context.Context) ([]domain.Role, error)
	New(ctx context.Context, name string) (*domain.Role, error)
	Remove(ctx context.Context, id int) (*domain.Id, error)
	Update(ctx context.Context, id int, name string) (*domain.Role, error)
}

