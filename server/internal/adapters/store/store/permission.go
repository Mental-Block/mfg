package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type PermissionStore struct {
	db *postgres.Store
}

func NewPermissionStore(db *postgres.Store) *PermissionStore {
	return &PermissionStore{
		db,
	}
}

func (pg *PermissionStore) Delete(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`DELETE FROM %v.permission WHERE permission_id = @id`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error())) 
	}

	return &id, nil
}

func (pg *PermissionStore) Insert(ctx context.Context, name string) (*domain.Permission, error) {
	query := fmt.Sprintf(`
		INSERT INTO %v.permission
		(
			name
		)
		VALUES 
		(
			@name
		)
		RETURNING 
			permission_id
			,name;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"name": name,
	}

	permisison := &domain.Permission{}
	err := pg.db.QueryRow(ctx, query, args).Scan(
		&permisison.Id,
		&permisison.Name,
	)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}

	return permisison, nil
}

func (pg *PermissionStore) Select(ctx context.Context, id domain.Id) (*domain.Permission, error) {
	query := fmt.Sprintf(`
		SELECT 
			permission_id
			,name
		FROM %s.permission
		WHERE permission_id = @id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var permision = &domain.Permission{}
	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&permision.Id,
		&permision.Name,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	return permision, nil
}

func (pg *PermissionStore) SelectPermissions(ctx context.Context) ([]domain.Permission, error) {
	query := fmt.Sprintf(`
		SELECT 
			permission_id as id
			,name
		FROM %s.permission
	`, postgres.PublicSchema)

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.Permission])
}


func (pg *PermissionStore) Update(ctx context.Context, id domain.Id, name string) (*domain.Permission, error) {
	query := fmt.Sprintf(`
		UPDATE %s.permission
		SET
			name=@name
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE permission_id = @id
		RETURNING 
			permission_id
			,name
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id":        id,
		"name":      name,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	permision := &domain.Permission{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&permision.Id,
		&permision.Name,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrPermissionNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to update row: %s", err.Error())) 
	}

	return permision, nil
}
