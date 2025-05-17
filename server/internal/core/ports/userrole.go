package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type UserRoleStore interface {
	// assigns a role to a user
	Assign(ctx context.Context, userId domain.Id, roleId domain.Id) error
	// unassigns a role to a user
	UnAssign(ctx context.Context, userId domain.Id, roleId domain.Id) error
}