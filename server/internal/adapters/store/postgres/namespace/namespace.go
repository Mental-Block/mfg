package namespace_store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/namespace/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of NameSpaceStore should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type INameSpaceStore interface {
	Select(ctx context.Context, id utils.UUID) (*NamespaceModel, error)
	Selects(ctx context.Context) ([]NamespaceModel, error)
	Update(ctx context.Context, input domain.Namespace) (*NamespaceModel, error)
	Upsert(ctx context.Context, input domain.Namespace) (*NamespaceModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteSoft(ctx context.Context, id utils.UUID) (*NamespaceModel, error)
}

type NamespaceStore struct {
	db *postgres.Store
}

func NewNamespaceStore(db *postgres.Store) *NamespaceStore {
	return &NamespaceStore{
		db,
	}
}

func (pg NamespaceStore) Select(ctx context.Context, id utils.UUID) (*NamespaceModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			namespace_id
			,name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE namespace_id = @id
		LIMIT 1;
	`,postgres.PublicSchema, postgres.TABLE_NAMESPACE)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &NamespaceModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,       
		&model.Name,     
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

func (pg NamespaceStore) Selects(ctx context.Context) ([]NamespaceModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			namespace_id
			,name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		LIMIT 100;
	`,postgres.PublicSchema, postgres.TABLE_NAMESPACE)

	rows, err := pg.db.Query(ctx, query)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []NamespaceModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[NamespaceModel])
}

func (pg NamespaceStore) Upsert(ctx context.Context, input domain.SanitizedNamespace) (*NamespaceModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Marshal")
	}

	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			namespace_id
			,name
			,metadata
		)
		VALUES 
		(
			@id
			,@name
			,@metadata
		)
		ON CONFLICT (namespace_id) DO UPDATE
		SET    
		(
			name
			,metadata
			,updated_dt
			,updated_by
		) = (
			 EXCLUDED.name
			,EXCLUDED.metadata
			,@updatedDT
			,@updatedBy
		)
		RETURNING 
			namespace_id
			,name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_NAMESPACE)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"name": input.Name,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &NamespaceModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,       
		&model.Name,     
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

func (pg NamespaceStore) Update(ctx context.Context, input domain.SanitizedNamespace) (*NamespaceModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Marshal")
	}

	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			username=@username	
			,active=@active
			,title=@title
			,avtar=@avatar
			,metadata=@metadata
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE namespace_id=@id
		RETURNING 
			namespace_id
			,name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_NAMESPACE)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"name": input.Name,
		"metadata": metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &NamespaceModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,       
		&model.Name,     
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

func (pg *NamespaceStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE namespace_id = @id;`, postgres.PublicSchema, postgres.TABLE_NAMESPACE)

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
 
func (pg *NamespaceStore) DeleteSoft(ctx context.Context, id utils.UUID) (*NamespaceModel, error) {
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			deleted_by=@updatedBy
			,deleted_dt=@updatedDT
		WHERE namespace_id = @id
		RETURNING 
			namespace_id
			,name
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_NAMESPACE)

	args := pgx.NamedArgs{
		"id":        id,
		"deletedBy": utils.NewDeletedBy(),
		"deletedDT": utils.NewDeletedDT(),
	}

	model := &NamespaceModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Name,
		&model.Metadata,
		&model.DeletedDT,
		&model.DeletedBy,
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
