package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/adapters/store/postgres/user"
	"github.com/server/internal/core/auth/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of AuthStore should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IAuthStore interface {
	AssignUser(ctx context.Context, id utils.UUID, userId utils.UUID) error
	Insert(ctx context.Context, input domain.SanitizedAuth) (*AuthModel, error) 
	Select(ctx context.Context, email domain.Email) (*AuthModel, error)
	SelectUser(ctx context.Context, id utils.UUID) (*user.UserModel, error)
	SelectVersion(ctx context.Context, id utils.UUID) (*int, error)
	SelectVersionByEmail(ctx context.Context, email domain.Email) (*int, error) 
	Update(ctx context.Context, input domain.SanitizedAuth) error
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
}

type AuthStore struct {
	db *postgres.Store
}

func NewAuthStore(db *postgres.Store) *AuthStore {
	return &AuthStore{
		db: db,
	}
}

func (pg* AuthStore) Insert(ctx context.Context, input domain.SanitizedAuth) (*AuthModel, error) {
	query := fmt.Sprintf(`
		INSERT INTO %v.%v
		(
			,auth_id
			,email
			,password
			,version
			,majic_active
			,otp_active
			,oidc_active
			,password_active
		)
		VALUES 
		(
			@id
			,@email
			,@password
			,@version
			,@majicActive
			,@otpActive
			,@oidcActive
			,@passwordActive
		)
		RETURNING 
			auth_id
			,email 
			,password
			,version
			,majic_active
			,otp_active
			,oidc_active
			,password_active
			,updated_by
			,updated_dt
			,created_by
			,created_dt;
	`, postgres.PublicSchema, postgres.TABLE_AUTH)

	args := pgx.NamedArgs{
		"id": 		 input.Id,
		"password":  input.Password,
		"email":     input.Email,
		"version":	 input.Version,
		"majicActive": input.MajicActive,	
		"otpActive": input.OTPActive,
		"oidcActive": input.OIDCActive,
		"passwordActive": input.PasswordActive,
	}

	model := &AuthModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Email,
		&model.Password,
		&model.Version,
		&model.MajicActive,
		&model.OTPActive,
		&model.OIDCActive,
		&model.PasswordActive,
		&model.UpdatedBy,
		&model.UpdatedDT,
		&model.CreatedBy,
		&model.CreatedDT,
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

func (pg* AuthStore) Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) {
	query := fmt.Sprintf(`DELETE FROM %v.%v WHERE auth_id = @id`, postgres.PublicSchema, postgres.TABLE_AUTH)

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

func (pg *AuthStore) SelectVersion(ctx context.Context, id utils.UUID) (*int, error) {
	query := fmt.Sprintf(`
		SELECT 
			version
		FROM %s.%v
		WHERE auth_id = @id
		LIMIT 1;
	`,postgres.PublicSchema, postgres.TABLE_AUTH)

	args := pgx.NamedArgs{
		"id": id,
	}

	var version *int

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&version,
	)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return version, nil
			default:
				return nil, err
		}
	}

	return version, nil
}

func (pg *AuthStore) SelectVersionByEmail(ctx context.Context, email domain.Email) (*int, error) {
	query := fmt.Sprintf(`
		SELECT 
			version
		FROM %s.%v
		WHERE email = @email
		LIMIT 1;
	`,postgres.PublicSchema, postgres.TABLE_AUTH)

	args := pgx.NamedArgs{
		"email": email,
	}

	var version *int

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&version,
	)

		if err != nil {
		err = postgres.CheckPostgresError(err)

		switch {
			case errors.Is(err, pgx.ErrNoRows):
				return version, nil
			default:
				return nil, err
		}
	}

	return version, nil
}

func (pg *AuthStore) SelectUser(ctx context.Context, id utils.UUID) (*user.UserModel, error) {
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
		FROM %s.%v
		INNER JOIN %s.%v USING (auth_id);
		INNER JOIN %s.%v as usr USING (user_id)
		WHERE auth_id = @id
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
		"id": id,
	}

	var model = &user.UserModel{}

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

func (pg *AuthStore) AssignUser(ctx context.Context, id utils.UUID, userId utils.UUID) error {
	query := fmt.Sprintf(`INSERT INTO %s.%v (auth_id, user_id) VALUES (auth_id, user_id)`,postgres.PublicSchema, postgres.TABLE_USER_AUTH)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

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

func (pg *AuthStore) Select(ctx context.Context, email domain.Email) (*AuthModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			auth_id
			,email 
			,password
			,version
			,majic_active
			,otp_active
			,oidc_active
			,password_active
			,updated_by
			,updated_dt
			,created_by
			,created_dt
		FROM %s.%v
		WHERE email = @email
		LIMIT 1;
	`,postgres.PublicSchema, postgres.TABLE_AUTH)

	args := pgx.NamedArgs{
		"email": email,
	}

	var model = &AuthModel{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&model.Id,
		&model.Email,
		&model.Password,
		&model.Version,
		&model.MajicActive,
		&model.OTPActive,
		&model.OIDCActive,
		&model.PasswordActive,
		&model.UpdatedBy,
		&model.UpdatedDT,
		&model.CreatedBy,
		&model.CreatedDT,
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

func (pg *AuthStore) Update(ctx context.Context, input domain.SanitizedAuth) error {
	query := fmt.Sprintf(`
		UPDATE %s.%v
		SET
			email=@email 
			,version=@version
			,password=@password
			,majic_active=@majicActive
			,otp_active=@otpActive
			,oidc_active=@oidcActive
			,password_active=@passwordActive
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE id = @id
	`,
	postgres.PublicSchema,
	postgres.TABLE_AUTH,
)

	args := pgx.NamedArgs{
		"password":  input.Password,
		"email":     input.Email,
		"version":   input.Version,
		"id": 		 input.Id,
		"majicActive": input.MajicActive,	
		"otpActive": input.OTPActive,
		"oidcActive": input.OIDCActive,
		"passwordActive": input.PasswordActive,
		"updatedBy": utils.NewDeletedBy(),
		"updatedDT": utils.NewDeletedDT(),
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		err = postgres.CheckPostgresError(err)

		return err
	}

	return nil
}
