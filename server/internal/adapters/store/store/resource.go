package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type ResourceStore struct {
	db *postgres.Store
}

func NewResourceStore(db *postgres.Store) *ResourceStore {
	return &ResourceStore{
		db,
	}
}

func (pg* ResourceStore) Delete(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`DELETE FROM %v.resource WHERE resource_id = @id`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error())) 
	}
	
	return &id, nil
}

// TODO - add atributes
func (pg* ResourceStore) Insert(ctx context.Context, name string) (*domain.Resource, error) {
	query := fmt.Sprintf(`
		INSERT INTO %v.resource
		(
			name
		)
		VALUES 
		(
			@name
		)
		RETURNING 
			resource_id
			,name;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"name": name,
	}

	resource := &domain.Resource{}
	err := pg.db.QueryRow(ctx, query, args).Scan(
		&resource.Id,
		&resource.Name,
	)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}

	return resource, nil
}

func (pg *ResourceStore) SelectResources(ctx context.Context) ([]domain.Resource, error) {
	query := fmt.Sprintf(`
		SELECT 
			resource_id as id
			,name
		FROM %s.resource
	`, postgres.PublicSchema)

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.Resource])
}

func (pg *ResourceStore) Select(ctx context.Context, id domain.Id) (*domain.Resource, error) {
	query := fmt.Sprintf(`
		SELECT 
			resource_id
			,name
		FROM %s.resource
		WHERE resource_id = @id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var resource = &domain.Resource{}
	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&resource.Id,
		&resource.Name,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	return resource, nil
}


/* TODO - add update attributes for resources */
func (pg *ResourceStore) Update(ctx context.Context, id domain.Id, name string) (*domain.Resource, error) {
	query := fmt.Sprintf(`
		UPDATE %s.resource
		SET
			name=@name
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE resource_id = @id
		RETURNING 
			resource_id
			,name
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id":        id,
		"name":      name,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	resource := &domain.Resource{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&resource.Id,
		&resource.Name,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to update row: %s", err.Error())) 
	}

	return resource, nil
}
