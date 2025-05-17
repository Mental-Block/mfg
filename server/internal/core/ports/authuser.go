package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type AuthUserStore interface {
	// gets the user roles, permissions, resources and attributes for user. (ABAC)
	Select(ctx context.Context, id domain.Id) (*domain.UserAuth, error)
}