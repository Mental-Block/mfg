package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type UserResourcePermissionStore struct {
	db *postgres.Store
}

func NewUserResourcePermissionStore(db *postgres.Store) *UserResourcePermissionStore {
	return &UserResourcePermissionStore{
		db,
	}
}

func (pg *UserResourcePermissionStore) Assign(ctx context.Context, userId domain.Id, resourcePermissionId domain.Id) error {
	query := fmt.Sprintf(`
		INSERT INTO %v.user_resource_permission
		(
			user_id
			,resource_permission_id
		)
		VALUES
		(
			@userId
			,@resourcePermissionId
		)
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"userId":  userId,
		"resourcePermissionId":  resourcePermissionId,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}

	return nil
}

func (pg *UserResourcePermissionStore) UnAssign(ctx context.Context, userId domain.Id, resourcePermissionId domain.Id) error {
	query := fmt.Sprintf(`DELETE FROM %v.user_resource_permission WHERE user_id = @userId AND resource_permission_id = @resourcePermissionId`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"userId":  userId,
		"resourcePermissionId":  resourcePermissionId,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserResourcePermissionNotFound.Error()) 
		}

		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error())) 
	}

	return nil
}
