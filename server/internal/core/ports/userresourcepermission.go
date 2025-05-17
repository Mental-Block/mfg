package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type UserResourcePermissionStore interface {
	// assigns a resource permission to a user
	Assign(ctx context.Context, userId domain.Id, resourcePermissionId domain.Id) error
	// unassigns a resource permission to a user
	UnAssign(ctx context.Context, userId domain.Id, resourcePermissionId domain.Id ) error
}