package role

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	role "github.com/server/internal/core/role/domain"
	"github.com/server/pkg/utils"
)

/*
	High level overview of RoleStore should not be directly imported.
	Copy interface and use dependancy injection over direct import.
*/
type IRoleStore interface {
	Upsert(ctx context.Context, input role.SanitizedRole) (*RoleModel, error)
	Update(ctx context.Context, input role.SanitizedRole) (*RoleModel, error) 
	Select(ctx context.Context, id utils.UUID) (*RoleModel, error)
	Selects(ctx context.Context) ([]RoleModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteSoft(ctx context.Context, id utils.UUID) (*RoleModel, error) 
}

type RoleStore struct {
	db *postgres.Store
}

func NewRoleStore(db *postgres.Store) *RoleStore {
	return &RoleStore{
		db: db,
	}
}

func (pg RoleStore) Upsert(ctx context.Context, input role.SanitizedRole) (*RoleModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to serialize metadata")
	}

	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			role_id
			,organization_id
			,name
			,title
			,permissions
			,active
			,scopes
			,metadata
		)
		VALUES 
		(
			@id
			,@active
			,@orgId
			,@name
			,@title
			,@permissions
			,@scopes
			,@metadata
		)
		ON CONFLICT (role_id) DO UPDATE
		SET    
		(
			organization_id
			,name
			,title
			,permissions
			,active
			,scopes
			,metadata
			,updated_by
			,updated_dt
		) = (
			EXCLUDED.organization_id
			,EXCLUDED.name
			,EXCLUDED.title
			,EXCLUDED.permissions
			,EXCLUDED.active
			,EXCLUDED.scopes
			,EXCLUDED.metadata
			,@updatedBy
			,@updatedDT
		)
		RETURNING 
			role_id
			,organization_id
			,name
			,title
			,permissions
			,active
			,scopes
			,metadata
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by;
	`, postgres.PublicSchema, postgres.TABLE_ROLE)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"active": input.Active,
		"orgId": input.OrgId,
		"name": input.Name,
		"title": input.Title,
		"permissions": input.Permissions,
		"scopes": input.Scopes,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &RoleModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.OrgId,
		&model.Name,
		&model.Title,
		&model.Permissions,
		&model.Active,
		&model.Scopes,
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

func (pg RoleStore) Update(ctx context.Context, input role.SanitizedRole) (*RoleModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to serialize metadata")
	}

	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			organization_id=@orgId
			,name=@name
			,title=@title
			,permissions=@permissions
			,active=@active
			,scopes=@scopes
			,metadata=@metadata
		WHERE role_id = @id
		RETURNING 
			role_id
			,organization_id
			,name
			,title
			,permissions
			,active
			,scopes
			,metadata
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by;
	`, postgres.PublicSchema, postgres.TABLE_ROLE)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"active": input.Active,
		"orgId": input.OrgId,
		"name": input.Name,
		"title": input.Title,
		"permissions": input.Permissions,
		"scopes": input.Scopes,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &RoleModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.OrgId,
		&model.Name,
		&model.Title,
		&model.Permissions,
		&model.Active,
		&model.Scopes,
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

func (pg RoleStore) Select(ctx context.Context, id utils.UUID) (*RoleModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			role_id
			,organization_id
			,name
			,title
			,permissions
			,active
			,scopes
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE role_id = @id
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &RoleModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.OrgId,
		&model.Name,
		&model.Title,
		&model.Permissions,
		&model.Active,
		&model.Scopes,
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

func (pg RoleStore) Selects(ctx context.Context) ([]RoleModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			role_id
			,organization_id
			,name
			,title
			,permissions
			,active
			,scopes
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		LIMIT 100;
	`, postgres.PublicSchema, postgres.TABLE_PERMISSION)

	rows, err := pg.db.Query(ctx, query)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []RoleModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[RoleModel])
}

func (pg RoleStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE role_id = @id;`, postgres.PublicSchema, postgres.TABLE_ROLE)

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
	
func (pg RoleStore) DeleteSoft(ctx context.Context, id utils.UUID) (*RoleModel, error) {
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			deleted_by=@deletedBy
			,deleted_dt=@deletedDT
		WHERE role_id = @id
		RETURNING 
			role_id
			,organization_id
			,name
			,title
			,permissions
			,active
			,scopes
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_ROLE)

	args := pgx.NamedArgs{
		"id":        id,
		"deletedBy": utils.NewDeletedBy(),
		"deletedDT": utils.NewDeletedDT(),
	}

	model := &RoleModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.OrgId,
		&model.Name,
		&model.Title,
		&model.Permissions,
		&model.Active,
		&model.Scopes,
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
	


