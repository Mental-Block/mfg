package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type RoleResourcePermissionStore struct {
	db *postgres.Store
}

func NewRoleResourcePermissionStore(db *postgres.Store) *RoleResourcePermissionStore {
	return &RoleResourcePermissionStore{
		db,
	}
}

func (pg *RoleResourcePermissionStore) Assign(ctx context.Context, roleId domain.Id, resourcePermissionId domain.Id) error {
	query := fmt.Sprintf(`
		INSERT INTO %v.role_resource_permission
		(
			role_id
			,resource_permission_id
		)
		VALUES
		(
			@roleId
			,@resourcePermissionId
		)
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"roleId":  roleId,
		"resourcePermissionId":  resourcePermissionId,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert role: %s", err.Error())) 
	}

	return nil
}

func (pg *RoleResourcePermissionStore) UnAssign(ctx context.Context, roleId domain.Id, resourcePermissionId domain.Id) error {
	query := fmt.Sprintf(`DELETE FROM %v.role_resource_permission WHERE role_id = @roleId AND resource_permission_id = @resourcePermissionId`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"resourcePermissionId":  resourcePermissionId,
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
