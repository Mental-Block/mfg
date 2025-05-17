package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type AuthUserStore struct {
	db *postgres.Store
}

func NewAuthUserStore(db *postgres.Store) *AuthUserStore {
	return &AuthUserStore{
		db,
	}
}

func (pg* AuthUserStore) Select(ctx context.Context, id domain.Id) (*domain.UserAuth, error) {
	query := fmt.Sprintf(`
		SELECT 
			usr.user_id
			,usr.username
		FROM %s.user AS usr
		WHERE usr.auth_id = @id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var user = &domain.UserAuth{
		Roles: []string{"resource:permission"},
	}

	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Username,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	return user, nil
}