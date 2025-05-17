package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type UserAuthStore interface {
	// removes authentication on the user table. authid becomes nil
	Remove(ctx context.Context, id domain.Id) (*domain.Id, error)
	// updates authentication on the user table. replaces authId
	Update(ctx context.Context, id domain.Id, authId domain.Id) error
}