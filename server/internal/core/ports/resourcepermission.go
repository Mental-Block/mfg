package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type ResourcePermissionStore interface {
	// creates a resource permission
	Insert(ctx context.Context, permissionId domain.Id, resourceId domain.Id) error
	// deletes a resource permission
	Delete(ctx context.Context, permissionResourceId domain.Id) error
}