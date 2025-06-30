package relation_store

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/relation/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of RealtionStore should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IRelationStore interface {
	Select(ctx context.Context, id utils.UUID) (*RelationModel, error)
	Selects(ctx context.Context) ([]RelationModel, error)
	SelectsByFields(ctx context.Context, input domain.SanitizedRelation) ([]RelationModel, error)
	Update(ctx context.Context, input domain.SanitizedRelation) (*RelationModel, error)
	Upsert(ctx context.Context, input domain.SanitizedRelation) (*RelationModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteSoft(ctx context.Context, id utils.UUID) (*RelationModel, error)
}

type RelationStore struct {
	db *postgres.Store
}

func NewRelationStore(db *postgres.Store) *RelationStore {
	return &RelationStore{
		db,
	}
}

func (pg RelationStore) Upsert(ctx context.Context, input domain.SanitizedRelation) (*RelationModel, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			relation_id
			,relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
		)
		VALUES 
		(
			@id
			,@relationName
			,@subjectId
			,@subjectNamespaceName
			,@subjectSubrelationName
			,@objectId
			,@objectNamespaceName
		)
		ON CONFLICT (relation_id) DO UPDATE
		SET    
		(
			relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,updated_by
			,updated_dt
		) = (
			EXCLUDED.relation_name
			,EXCLUDED.subject_id
			,EXCLUDED.subject_namespace_name
			,EXCLUDED.subject_subrelation_name
			,EXCLUDED.object_id
			,EXCLUDED.object_namespace_name
			,@updatedBy
			,@updatedDT
		)
		RETURNING 
			relation_id
			,relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by;
	`, postgres.PublicSchema, postgres.TABLE_RELATION)
	
	args := pgx.NamedArgs{
		"id": input.Id,
		"relationName": input.RelationName,
		"subjectId": input.Subject.Id,
		"subjectNamespaceName": input.Subject.Namespace,
		"subjectSubrelationName": input.Subject.SubRelationName,
		"objectId": input.Object.Id,
		"objectNamespaceName": input.Object.Namespace,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &RelationModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.RelationName,
		&model.SubjectId,
		&model.SubjectNamespace,
		&model.SubjectSubRelation,
		&model.ObjectId,
		&model.ObjectNamespace,
		&model.CreatedDT,
		&model.CreatedBy,
		&model.UpdatedDT,
		&model.UpdatedBy,
		&model.DeletedDT,
		&model.DeletedBy,
	)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		return nil, err
	}

	return model, nil
}

func (pg RelationStore) Update(ctx context.Context, input domain.SanitizedRelation) (*RelationModel, error) {
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,updated_by
			,updated_dt
		WHERE relation_id = @id
		RETURNING 
			relation_id
			,relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_RELATION)

	args := pgx.NamedArgs{
		"id": input.Id,
		"relationName": input.RelationName,
		"subjectId": input.Subject.Id,
		"subjectNamespaceName": input.Subject.Namespace,
		"subjectSubrelationName": input.Subject.SubRelationName,
		"objectId": input.Object.Id,
		"objectNamespaceName": input.Object.Namespace,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &RelationModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.RelationName,
		&model.SubjectId,
		&model.SubjectNamespace,
		&model.SubjectSubRelation,
		&model.ObjectId,
		&model.ObjectNamespace,
		&model.CreatedDT,
		&model.CreatedBy,
		&model.UpdatedDT,
		&model.UpdatedBy,
		&model.DeletedDT,
		&model.DeletedBy,
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

func (pg RelationStore) Select(ctx context.Context, id utils.UUID) (*RelationModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			relation_id
			,relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by
		FROM %s.%s
		WHERE relation_id = @id
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_RELATION)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &RelationModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.RelationName,
		&model.SubjectId,
		&model.SubjectNamespace,
		&model.SubjectSubRelation,
		&model.ObjectId,
		&model.ObjectNamespace,
		&model.CreatedDT,
		&model.CreatedBy,
		&model.UpdatedDT,
		&model.UpdatedBy,
		&model.DeletedDT,
		&model.DeletedBy,
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

func (pg RelationStore) Selects(ctx context.Context) ([]RelationModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			relation_id
			,relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by
		FROM %s.%s
		WHERE relation_id = @id
		LIMIT 100;
	`, postgres.PublicSchema, postgres.TABLE_RELATION)

	rows, err := pg.db.Query(ctx, query)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []RelationModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[RelationModel])
}

func (pg RelationStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE relation_id = @id;`, postgres.PublicSchema, postgres.TABLE_RELATION)

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
	
func (pg RelationStore) DeleteSoft(ctx context.Context, id utils.UUID) (*RelationModel, error) {
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			,deleted_by=@deletedBy
			,deleted_dt=@deletedDT
		WHERE relation_id = @id
		RETURNING 
			relation_id
			,relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_RELATION)

	args := pgx.NamedArgs{
		"id":        id,
		"deletedBy": utils.NewUpdatedBy(),
		"deletedDT": utils.NewUpdatedDT(),
	}

	model := &RelationModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.RelationName,
		&model.SubjectId,
		&model.SubjectNamespace,
		&model.SubjectSubRelation,
		&model.ObjectId,
		&model.ObjectNamespace,
		&model.CreatedDT,
		&model.CreatedBy,
		&model.UpdatedDT,
		&model.UpdatedBy,
		&model.DeletedDT,
		&model.DeletedBy,
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
		

func (pg RelationStore) SelectsByFields(ctx context.Context, input domain.SanitizedRelation) ([]RelationModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			relation_id
			,relation_name
			,subject_id
			,subject_namespace_name
			,subject_subrelation_name
			,object_id
			,object_namespace_name
			,created_dt
			,created_by
			,updated_dt
			,updated_by
			,deleted_dt
			,deleted_by
		FROM %s.%s
		WHERE
			(@objectId IS NULL OR object_id = @objectId) OR
			(@ObjectNameSpaceName IS NULL OR object_namespace_name = @ObjectNameSpaceName) OR
			(@subjectId IS NULL OR subject_id = @subjectId) OR
			(@subjectNameSpaceName IS NULL OR subject_namespace_name = @subjectNamespaceName) OR
			(@relationName IS NULL OR relation_name = @relationName);
	`, postgres.PublicSchema, postgres.TABLE_RELATION)

	args := pgx.NamedArgs{ 
		"objectId": input.Object.Id.String(),
		"ObjectNameSpaceName": input.Object.Namespace.String(),
		"subjectId": input.Subject.Id,
		"subjectNameSpaceName": input.Subject.Namespace.String(),
		"relationName": input.RelationName.String(),
	}
	
	rows, err := pg.db.Query(ctx, query, args)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []RelationModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[RelationModel])
}
