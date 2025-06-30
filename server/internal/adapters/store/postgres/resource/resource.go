package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/resource/domain"
	"github.com/server/pkg/utils"
)

/*
	High level overview of ResourceStore should not be directly imported.
	Copy interface and use dependancy injection over direct import.
*/
type IResourceStore interface {
	Upsert(ctx context.Context, input domain.SanitizedResource) (*ResourceModel, error)
	Update(ctx context.Context, input domain.SanitizedResource) (*ResourceModel, error) 
	Select(ctx context.Context, id utils.UUID) (*ResourceModel, error)
	SelectByURN(ctx context.Context, urn string) (*ResourceModel, error)
	Selects(ctx context.Context) ([]ResourceModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) 
	DeleteSoft(ctx context.Context, id utils.UUID) (*ResourceModel, error) 
}

type ResourceStore struct {
	db *postgres.Store
}

func NewResourceStore(db *postgres.Store) *ResourceStore {
	return &ResourceStore{
		db: db,
	}
}

func (pg ResourceStore) Upsert(ctx context.Context, input domain.SanitizedResource) (*ResourceModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to serialize metadata")
	}
	
	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			resource_id
			,urn
			,name
			,title
			,project_id
			,namespace_name
			,principal_id
			,principal_type
			,metadata
		)
		VALUES 
		(
			@id
			,@urn
			,@name
			,@title
			,@projectId
			,@namespaceId
			,@principalId
			,@principalType
			,@metadata
		)
		ON CONFLICT (resource_id) DO UPDATE
		SET    
		(
			urn
			,name
			,title
			,project_id
			,namespace_name
			,principal_id
			,principal_type
			,metadata
		) = (
			EXCLUDED.urn
			,EXCLUDED.name
			,EXCLUDED.title
			,EXCLUDED.project_id
			,EXCLUDED.namespace_name
			,EXCLUDED.principal_id
			,EXCLUDED.principal_type
			,EXCLUDED.metadata
			,@updatedBy
			,@updatedDT
		)
		RETURNING 
			resource_id
			,urn
			,name
			,title
			,project_id
			,namespace_name
			,principal_id
			,principal_type
			,metadata
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by;
	`, postgres.PublicSchema, postgres.TABLE_RESOURCE)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"urn":input.URN,
		"name":input.Name,
		"title":input.Title,
		"projectId":input.ProjectId,
		"namespaceId":input.NamespaceId,
		"principalId":input.PrincipalId,
		"principalType":input.PrincipalType,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &ResourceModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.URN,
		&model.Name,
		&model.Title,
		&model.ProjectId,
		&model.NamespaceId,
		&model.PrincipalId,
		&model.PrincipalType,
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

func (pg ResourceStore) Update(ctx context.Context, input domain.SanitizedResource) (*ResourceModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to serialize metadata")
	}
	
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			urn=@urn
			,name=@name
			,title=@title
			,project_id=@projectId
			,namespace_name=@namespaceId
			,principal_id=@principalId
			,principal_type=@principalType
			,metadata=@metadata
			,updated_dt=@updatedDT
			,updated_by=@updatedBy
		WHERE resource_id = @id
		RETURNING 
			urn
			,name
			,title
			,project_id
			,namespace_name
			,principal_id
			,principal_type
			,metadata
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by;
	`, postgres.PublicSchema, postgres.TABLE_RESOURCE)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"urn":input.URN,
		"name":input.Name,
		"title":input.Title,
		"projectId":input.ProjectId,
		"namespaceId":input.NamespaceId,
		"principalId":input.PrincipalId,
		"principalType":input.PrincipalType,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &ResourceModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.URN,
		&model.Name,
		&model.Title,
		&model.ProjectId,
		&model.NamespaceId,
		&model.PrincipalId,
		&model.PrincipalType,
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

func (pg ResourceStore) Select(ctx context.Context, id utils.UUID) (*ResourceModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			resource_id
			,urn
			,name
			,title
			,project_id
			,namespace_name
			,principal_id
			,principal_type
			,metadata
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by
		FROM %s.%s
		WHERE resource_id = @id
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_RESOURCE)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &ResourceModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.URN,
		&model.Name,
		&model.Title,
		&model.ProjectId,
		&model.NamespaceId,
		&model.PrincipalId,
		&model.PrincipalType,
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

func (pg ResourceStore) Selects(ctx context.Context) ([]ResourceModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			resource_id
			,urn
			,name
			,title
			,project_id
			,namespace_name
			,principal_id
			,principal_type
			,metadata
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by
		FROM %s.%s
		LIMIT 100;
	`, postgres.PublicSchema, postgres.TABLE_RESOURCE)

	rows, err := pg.db.Query(ctx, query)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []ResourceModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[ResourceModel])
}

func (pg ResourceStore) SelectByURN(ctx context.Context, urn string) (*ResourceModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			resource_id
			,urn
			,name
			,title
			,project_id
			,namespace_name
			,principal_id
			,principal_type
			,metadata
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by
		FROM %s.%s
		WHERE urn = @urn
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_RESOURCE)

	args := pgx.NamedArgs{
		"urn": urn,
	}

	model := &ResourceModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.URN,
		&model.Name,
		&model.Title,
		&model.ProjectId,
		&model.NamespaceId,
		&model.PrincipalId,
		&model.PrincipalType,
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

func (pg ResourceStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE resource_id = @id;`, postgres.PublicSchema, postgres.TABLE_RESOURCE)

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


func (pg ResourceStore) DeleteSoft(ctx context.Context, id utils.UUID) (*ResourceModel, error) {
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			,deleted_by=@deletedBy
			,deleted_dt=@deletedDT
		WHERE role_id = @id
		RETURNING 
			role_id
			,org_id
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
	`, postgres.PublicSchema, postgres.TABLE_RESOURCE)

	args := pgx.NamedArgs{
		"id":        id,
		"deletedBy": utils.NewDeletedBy(),
		"deletedDT": utils.NewDeletedDT(),
	}

	model := &ResourceModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.URN,
		&model.Name,
		&model.Title,
		&model.ProjectId,
		&model.NamespaceId,
		&model.PrincipalId,
		&model.PrincipalType,
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
	
