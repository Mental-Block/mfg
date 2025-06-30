package permission

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/permission/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of PermissionStore should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IPermissionStore interface {
	Select(ctx context.Context, id utils.UUID) (*PermissionModel, error)
	SelectBySlug(ctx context.Context, slug string) (*PermissionModel, error)
	Selects(ctx context.Context) ([]PermissionModel, error)
	Upsert(ctx context.Context, input domain.SanitizedPermission) (*PermissionModel, error)
	Update(ctx context.Context, input domain.SanitizedPermission) (*PermissionModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) 
	DeleteSoft(ctx context.Context, id utils.UUID) (*PermissionModel, error)
}

type PermissionStore struct {
	db *postgres.Store
}

func NewPermissionStore(db *postgres.Store) *PermissionStore {
	return &PermissionStore{
		db: db,
	}
}

func (pg *PermissionStore) Upsert(ctx context.Context, input domain.SanitizedPermission) (*PermissionModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Marshal")
	}

	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			permission_id
			,name
			,slug
			,namespace_name
			,metadata
		)
		VALUES 
		(
			@id
			,@name
			,@slug
			,@namespaceId
			,@metadata
		)
		ON CONFLICT (permission_id) DO UPDATE
		SET    
		(
			name
			,slug
			,namespace_name
			,metadata
			,updated_by
			,updated_dt
		) = (
			EXCLUDED.name
			,EXCLUDED.slug
			,EXCLUDED.namespace_name
			,EXCLUDED.metadata
			,@updatedBy
			,@updatedDT
		)
		RETURNING 
			permission_id
			,name
			,slug
			,namespace_name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_RELATION)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"name": input.Name,
		"slug": input.Slug,
		"namespaceId": input.NamespaceId,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &PermissionModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Name,
		&model.Slug,
		&model.NamespaceId,
		&model.Metadata,
		&model.DeletedBy,
		&model.DeletedDT,
		&model.UpdateBy,
		&model.UpdatedDT,
		&model.CreatedBy,
		&model.CreatedDt,
	)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		return nil, err
	}

	return model, nil
}

func (pg PermissionStore) Update(ctx context.Context, input domain.SanitizedPermission) (*PermissionModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Marshal")
	}
	
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			name=@name
			,slug=@slug
			,namespace_name = @namespaceId
			,metadata = @metadata
			,updated_by = @updatedBy
			,updated_dt = @updatedDT
		WHERE permision_id = @id
		RETURNING 
			permission_id
			,name
			,slug
			,namespace_name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt; 
	`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	args := pgx.NamedArgs{
		"id": input.Id,
		"name": input.Name,
		"slug": input.Slug,
		"namespaceId": input.NamespaceId,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &PermissionModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Name,
		&model.Slug,
		&model.NamespaceId,
		&model.Metadata,
		&model.DeletedBy,
		&model.DeletedDT,
		&model.UpdateBy,
		&model.UpdatedDT,
		&model.CreatedBy,
		&model.CreatedDt,
	)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return model, nil
			default:
				return nil, err
		}
	}

	return model, nil
}

func (pg PermissionStore) Select(ctx context.Context, id utils.UUID) (*PermissionModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			permission_id
			,name
			,slug
			,namespace_name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE permision_id = @id
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &PermissionModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Name,
		&model.Slug,
		&model.NamespaceId,
		&model.Metadata,
		&model.DeletedBy,
		&model.DeletedDT,
		&model.UpdateBy,
		&model.UpdatedDT,
		&model.CreatedBy,
		&model.CreatedDt,
	)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return model, nil
			default:
				return nil, err
		}
	}

	return model, nil
}

func (pg PermissionStore) SelectBySlug(ctx context.Context, slug string) (*PermissionModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			permission_id
			,name
			,slug
			,namespace_name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE slug = @slug
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	args := pgx.NamedArgs{ 
		"slug": slug,
	}

	model := &PermissionModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Name,
		&model.Slug,
		&model.NamespaceId,
		&model.Metadata,
		&model.DeletedBy,
		&model.DeletedDT,
		&model.UpdateBy,
		&model.UpdatedDT,
		&model.CreatedBy,
		&model.CreatedDt,
	)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return model, nil
			default:
				return nil, err
		}
	}

	return model, nil
}

func (pg PermissionStore) Selects(ctx context.Context) ([]PermissionModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			permission_id
			,name
			,slug
			,namespace_name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE permission_id = @id
		LIMIT 100;
	`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	rows, err := pg.db.Query(ctx, query)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []PermissionModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[PermissionModel])
}

func (pg PermissionStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE permission_id = @id;`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return &id, nil
			default:
				return nil, err
		}
	}

	return &id, nil
}
	
func (pg PermissionStore) DeleteSoft(ctx context.Context, id utils.UUID) (*PermissionModel, error) {
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			deleted_by=@updatedBy
			,deleted_dt=@updatedDT
		WHERE permission_id = @id
		RETURNING 
			permission_id
			,name
			,slug
			,namespace_name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	args := pgx.NamedArgs{
		"id":        id,
		"deletedBy": utils.NewDeletedBy(),
		"deletedDT": utils.NewDeletedDT(),
	}

	model := &PermissionModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Name,
		&model.Slug,
		&model.NamespaceId,
		&model.Metadata,
		&model.DeletedBy,
		&model.DeletedDT,
		&model.UpdateBy,
		&model.UpdatedDT,
		&model.CreatedBy,
		&model.CreatedDt,
	)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return model, nil
			default:
				return nil, err
		}
	}

	return model, nil
}
