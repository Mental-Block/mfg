package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type RoleStore struct {
	db *postgres.Store
}

func NewRoleStore(db *postgres.Store) *RoleStore {
	return &RoleStore{
		db,
	}
}

func (pg* RoleStore) Delete(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`DELETE FROM %v.role WHERE role_id = @id`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error())) 
	}

	return &id, nil
}

func (pg* RoleStore) Insert(ctx context.Context, name string) (*domain.Role, error) {
	query := fmt.Sprintf(`
		INSERT INTO %v.role
		(
			name
		)
		VALUES 
		(
			@name
		)
		RETURNING 
			role_id
			,name;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"name": name,
	}

	role := &domain.Role{}
	err := pg.db.QueryRow(ctx, query, args).Scan(
		&role.Id,
		&role.Name,
	)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}

	return role, nil
}

func (pg *RoleStore) Select(ctx context.Context, id domain.Id) (*domain.Role, error) {
	query := fmt.Sprintf(`
		SELECT 
			role_id
			,name
		FROM %s.role
		WHERE role_id = @id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var role = &domain.Role{}
	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&role.Id,
		&role.Name,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	return role, nil
}

func (pg *RoleStore) SelectRoles(ctx context.Context) ([]domain.Role, error) {
	query := fmt.Sprintf(`
		SELECT 
			role_id as id
			,name
		FROM %s.role
	`, postgres.PublicSchema)

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.Role])
}

func (pg *RoleStore) Update(ctx context.Context, id domain.Id, name string) (*domain.Role, error) {
	query := fmt.Sprintf(`
		UPDATE %s.role
		SET
			name=@name
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE role_id = @id
		RETURNING 
			role_id
			,name
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id":        id,
		"name":      name,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	role := &domain.Role{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&role.Id,
		&role.Name,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrRoleNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to update row: %s", err.Error())) 
	}

	return role, nil
}
