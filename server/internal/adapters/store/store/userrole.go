package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type UserRoleStore struct {
	db *postgres.Store
}

func NewUserRoleStore(db *postgres.Store) *UserRoleStore {
	return &UserRoleStore{
		db,
	}
}

func (pg *UserRoleStore) Assign(ctx context.Context, userId domain.Id, roleId domain.Id) error {
	query := fmt.Sprintf(`
		INSERT INTO %v.user_role
		(
			user_id
			,role_id
		)
		VALUES
		(
			@userId
			,@roleId
		)
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"userId":  userId,
		"roleId":  roleId,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}

	return nil
}

func (pg *UserRoleStore) UnAssign(ctx context.Context, userId domain.Id, roleId domain.Id) error {
	query := fmt.Sprintf(`DELETE FROM %v.user_role WHERE user_id = @userId AND role_id = @roleId`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"userId":  userId,
		"roleId":  roleId,
	}

	_, err := pg.db.Exec(ctx, query, args)

	
	if err != nil {
		if err == pgx.ErrNoRows {
			return internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error()) 
		}

		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error())) 
	}


	return nil
}

