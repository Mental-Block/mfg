package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type RoleResourcePermissionStore interface {
	// assigns a resource permission to a user
	Assign(ctx context.Context, roleId domain.Id, resourcePermissionId domain.Id) error
	// unassigns a resource permission to a user
	UnAssign(ctx context.Context, roleId domain.Id, resourcePermissionId domain.Id ) error
}