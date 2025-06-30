package serviceuser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/serviceuser/domain"
	"github.com/server/pkg/utils"
)

/*
	High level overview of IServiceUserCredentialStore should not be directly imported.
	Copy interface and use dependancy injection over direct import.
*/
type IServiceUserCredentialStore interface {
	Select(ctx context.Context, id utils.UUID)(*ServiceUserCredentialModel, error)
	Selects(ctx context.Context, filter domain.CredentialFilter)([]ServiceUserCredentialModel, error)
 	Insert(ctx context.Context, input domain.SanitizedServiceUserCredential)(*ServiceUserCredentialModel, error)
	DeleteByUserId(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
}

type ServiceUserCredentialStore struct {
	db *postgres.Store
}

func NewServiceUserCredentialStore(db *postgres.Store) *ServiceUserCredentialStore {
	return &ServiceUserCredentialStore{
		db,
	}
}

func (pg ServiceUserCredentialStore) Select(ctx context.Context, id utils.UUID)(*ServiceUserCredentialModel, error){
	query := fmt.Sprintf(`
		SELECT 
			serviceuser_credential_id
			,serviceuser_id
			,type
			,secret_hash
			,public_key
			,title
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE serviceuser_credential_id = @id
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER_CREDENTIAL)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &ServiceUserCredentialModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.ServiceUserId,
		&model.Type,
		&model.SecretHash,
		&model.PublicKey,
		&model.Title,
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

func (pg ServiceUserCredentialStore) Selects(ctx context.Context, filter domain.CredentialFilter)([]ServiceUserCredentialModel, error){
	
	// do a better job filtering later with tsvector
	query := fmt.Sprintf(`
		SELECT 
			serviceuser_credential_id
			,serviceuser_id
			,type
			,secret_hash
			,public_key
			,title
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
			%s.%s AS su USING (serviceuser_id) 
		WHERE 
			(@id IS NULL OR serviceuser_credential_id = @id) OR
			(@orgId IS NULL OR su.organization_id = @ordId) OR
			(@serviceUserId IS NULL OR su.serviceuser_id = @serviceUserId) OR
			(@isPublicKey IS NULL OR (@isPublicKey IS TRUE secret_hash IS NOT NULL AND type = "jwt_bearer"))
			(@isToken IS NULL OR (@isToken IS TRUE AND secret_hash IS NOT NULL AND type = 'opaque_token'))
			(@isSecret IS NULL OR @isSecret IS TRUE AND secret_hash IS NOT NULL AND type = 'client_credential')
		LIMIT 100;
	`, 
		postgres.PublicSchema, 
		postgres.TABLE_SERVICEUSER_CREDENTIAL, 
		postgres.PublicSchema, 
		postgres.TABLE_SERVICEUSER,
	)
	
	args := pgx.NamedArgs{
		"id": filter.Id,
		"orgId": filter.OrgId,
		"serviceUserId": filter.ServiceUserId,
		"isPublicKey": filter.IsKey,       
		"isSecret": filter.IsSecret,
		"isToken": filter.IsToken,
	}

	rows, err := pg.db.Query(ctx, query, args)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []ServiceUserCredentialModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[ServiceUserCredentialModel])
}

func (pg ServiceUserCredentialStore) Insert(ctx context.Context, input domain.SanitizedServiceUserCredential)(*ServiceUserCredentialModel, error){
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Marshal")
	}

	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			serviceuser_id
			,type
			,secret_hash
			,public_key
			,title
			,metadata
		)
		VALUES 
		(
			@serviceuserId
			@type
			@secretHash
			@publicKey
			@title
			@metadata
		)
		RETURNING 
			serviceuser_credential_id
			,serviceuser_id
			,type
			,secret_hash
			,public_key
			,title
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER_CREDENTIAL)

	args := pgx.NamedArgs{
		"serviceuserId": input.ServiceUserId,
		"type":	input.Type,
		"secretHash":input.SecretHash,
		"publicKey": input.PublicKey,
		"title": input.Title,
		"metadata": metaData,
	}

	model := &ServiceUserCredentialModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.ServiceUserId,
		&model.Type,
		&model.SecretHash,
		&model.PublicKey,
		&model.Title,
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

func (pg ServiceUserCredentialStore) DeleteByUserId(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE serviceuser_id = @id;`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER_CREDENTIAL)

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

func (pg ServiceUserCredentialStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE serviceuser_credential_id = @id;`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER_CREDENTIAL)

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
