package serviceuser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/adapters/store/postgres/tx"
	"github.com/server/internal/core/serviceuser/domain"
	"github.com/server/pkg/utils"
)

/*
	High level overview of serviceUserStore should not be directly imported.
	Copy interface and use dependancy injection over direct import.
*/
type IServiceUserStore interface {
	Insert(ctx context.Context, input domain.SanitizedServiceUser) (*ServiceUserModel, error)
	Selects(ctx context.Context, filter domain.ServiceUserFilter) ([]ServiceUserModel, error) 
	SelectByIds(ctx context.Context, ids []utils.UUID) ([]ServiceUserModel, error)
	Select(ctx context.Context, id utils.UUID) (*ServiceUserModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteUserAndCredentials(ctx context.Context, id utils.UUID) error
}

type ServiceUserStore struct {
	db *postgres.Store
}

func NewServiceUserStore(db *postgres.Store) *ServiceUserStore {
	return &ServiceUserStore{
		db: db,
	}
}

func (pg ServiceUserStore) Insert(ctx context.Context, input domain.SanitizedServiceUser) (*ServiceUserModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Marshal")
	}

	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			serviceuser_id
			,organization_id
			,title
			,state
			,metadata
		)
		VALUES 
		(
			@serviceuserId
			@organizationId
			@title
			@state
			@metadata
		)
		RETURNING 
			serviceuser_id
			,organization_id
			,title
			,state
			,metadata
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER)

	args := pgx.NamedArgs{
		"serviceuserId": input.Id,
		"organizationId": input.OrgId,
		"title": input.Title,
		"state": input.State,
		"metadata": metaData,
	}

	model := &ServiceUserModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.OrganizationId,
		&model.Title,
		&model.State,
		&model.Metadata,
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


func (pg ServiceUserStore) Selects(ctx context.Context, filter domain.ServiceUserFilter) ([]ServiceUserModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			serviceuser_id
			,organization_id
			,title
			,state
			,metadata
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE
			(@serviceUserIds IS NULL OR serviceuser_id IN(@serviceUserId)) OR
			(@orgId IS NULL OR organization_id = @orgId) OR
			(@state IS NULL OR state = @state)
		LIMIT 100;
	`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER)

	args := pgx.NamedArgs{
		"serviceUserIds": strings.Join(filter.Ids, ","),
		"orgId": filter.OrgId, 
		"state": filter.State,
	}

	rows, err := pg.db.Query(ctx, query, args)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []ServiceUserModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[ServiceUserModel])
}

// returns a list of service users by their Ids.
func (pg ServiceUserStore) SelectByIds(ctx context.Context, uuids []utils.UUID) ([]ServiceUserModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			serviceuser_id
			,organization_id
			,title
			,state
			,metadata
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE 
			serviceuser_id IN (@ids);
	`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER)

	ids := make([]string, len(uuids))

	for i, _ := range ids {
		ids[i] = uuids[i].String()
	}

	args := pgx.NamedArgs {
		"ids": strings.Join(ids, ","),
	}

	rows, err := pg.db.Query(ctx, query, args)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []ServiceUserModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[ServiceUserModel])
}

func (pg ServiceUserStore) Select(ctx context.Context, id utils.UUID) (*ServiceUserModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			serviceuser_id
			,organization_id
			,title
			,state
			,metadata
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE serviceuser_id = @id
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &ServiceUserModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.OrganizationId,
		&model.Title,
		&model.State,
		&model.Metadata,
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

func (pg ServiceUserStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE serviceuser_id = @id;`, postgres.PublicSchema, postgres.TABLE_SERVICEUSER)

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


func (pg ServiceUserStore) DeleteUserAndCredentials(ctx context.Context, id utils.UUID) error {
	
	err := tx.RunInTx(ctx, pg.db.Pool, func(tx pgx.Tx) error {
		credentials := ServiceUserCredentialStore{
			db: pg.db,
		}

		_, err := credentials.DeleteByUserId(ctx, id)

		if (err != nil) {
			return err
		}

		_, err = pg.Delete(ctx, id)

		if err != nil {
			return err
		}
		
		return nil
	})

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return nil
			default:
				return err
		}
	}

	return nil
}