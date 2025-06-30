package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/server/internal/adapters/store/postgres"

	auth_domain "github.com/server/internal/core/auth/domain"
	"github.com/server/internal/core/user/domain"

	"github.com/server/pkg/utils"
)

/*
 High level overview of UserStore should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IUserStore interface {
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteSoft(ctx context.Context, id utils.UUID) (*UserModel, error) 
	Select(ctx context.Context, id utils.UUID) (*UserModel, error)
	Insert(ctx context.Context, input domain.SanitizedUser) (*UserModel, error) 
	Selects(ctx context.Context) ([]UserModel, error)
	SelectByEmail(ctx context.Context, email auth_domain.Email) (*UserModel, error) 
	Update(ctx context.Context, input domain.SanitizedUser) (*UserModel, error)
}

type UserStore struct {
	db *postgres.Store
}

func NewUserStore(db *postgres.Store) *UserStore {
	return &UserStore{
		db,
	}
}

func (pg *UserStore) DeleteSoft(ctx context.Context, id utils.UUID) (*UserModel, error) {
	query := fmt.Sprintf(`
		UPDATE %s.%s
		SET
			,deleted_by=@updatedBy
			,deleted_dt=@updatedDT
		WHERE user_id = @id
		RETURNING 
			user_id
			,username	
			,active
			,title
			,avatar
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_USER)

	args := pgx.NamedArgs{
		"id":        id,
		"deletedBy": utils.NewDeletedBy(),
		"deletedDT": utils.NewDeletedDT(),
	}

	model := &UserModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Username,
		&model.Active,
		&model.Title,
		&model.Avatar, 
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


func (pg *UserStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %s.%s WHERE user_id = @id;`, postgres.PublicSchema, postgres.TABLE_USER)

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

func (pg *UserStore) Select(ctx context.Context, id utils.UUID) (*UserModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			user_id
			,username	
			,active
			,title
			,avatar
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		WHERE user_id = @id
		LIMIT 1;
	`, postgres.PublicSchema, postgres.TABLE_USER)

	args := pgx.NamedArgs{
		"id": id,
	}

	model := &UserModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Username,
		&model.Active,
		&model.Title,
		&model.Avatar, 
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

func (pg *UserStore) Selects(ctx context.Context) ([]UserModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			user_id
			,username	
			,active
			,title
			,avatar
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%s
		LIMIT 100;
	`, postgres.PublicSchema, postgres.TABLE_USER)

	rows, err := pg.db.Query(ctx, query)

	defer rows.Close()

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return []UserModel{}, nil
			default:
				return nil, err
		}
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[UserModel])
}
   
func (pg *UserStore) Insert(ctx context.Context, input domain.SanitizedUser) (*UserModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "json.Marshal")
	}
	
	query := fmt.Sprintf(`
		INSERT INTO %s.%s
		(
			,user_id
			,username	
			,active
			,title
			,avatar
			,metadata
		)
		VALUES 
		(
			,@id
			,@username	
			,@active
			,@title
			,@avatar
			,@metadata
		)
		RETURNING 
			user_id
			,username	
			,active
			,title
			,avatar
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt
	`, postgres.PublicSchema, postgres.TABLE_USER)

	args := pgx.NamedArgs{
		"id": input.Id,
		"username": input.Username,
		"active": input.Active,
		"title": input.Title,
		"avtar": input.Avatar,
		"metadata": metaData,
	}

	model := &UserModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Username,
		&model.Active,
		&model.Title,
		&model.Avatar, 
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

func (pg *UserStore) Update(ctx context.Context, input domain.SanitizedUser) (*UserModel, error) {
	metaData, err := json.Marshal(input.Metadata)

	if err != nil {
		return  nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to serialize metadata")
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
		WHERE user_id = @id
		RETURNING 
			user_id
			,username	
			,active
			,title
			,avatar
			,metadata
			,deleted_by
			,deleted_dt
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_USER)

	args := pgx.NamedArgs{
		"id":        input.Id,
		"username":  input.Username,
		"active":    input.Active,
		"title": 	 input.Title,
		"avtar": 	 input.Avatar,
		"metadata":  metaData,
		"updatedBy": utils.NewUpdatedBy(),
		"updatedDT": utils.NewUpdatedDT(),
	}

	model := &UserModel{}

	err = pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Username,
		&model.Active,
		&model.Title,
		&model.Avatar, 
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

func (pg *UserStore) SelectByEmail(ctx context.Context, email auth_domain.Email) (*UserModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			usr.user_id
			,usr.username	
			,usr.active
			,usr.title
			,usr.avatar
			,usr.metadata
			,usr.deleted_by
			,usr.deleted_dt
			,usr.updated_by
			,usr.updated_dt
			,usr.created_by
			,usr.created_dt
		FROM %s.%s 
			INNER JOIN %s.%s USING (auth_id) 
			INNER JOIN %s.%s AS usr USING (user_id)
		WHERE auth.email = @email
		LIMIT 1;
	`, 	
	postgres.PublicSchema, 
	postgres.TABLE_AUTH, 
	postgres.PublicSchema, 
	postgres.TABLE_USER_AUTH, 
	postgres.PublicSchema,
	postgres.TABLE_USER,
)

	args := pgx.NamedArgs{
		"email": email,
	}

	var model = &UserModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Username,
		&model.Active,
		&model.Title,
		&model.Avatar, 
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


