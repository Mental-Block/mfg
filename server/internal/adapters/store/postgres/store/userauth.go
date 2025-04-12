package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type UserAuthStore struct {
	db *postgres.Store
}

func NewUserAuthStore(db *postgres.Store) *UserAuthStore {
	return &UserAuthStore{
		db,
	}
}

func (pg *UserAuthStore) UpdateVerified(ctx context.Context, email domain.Email) error {
	query := fmt.Sprintf(`
		UPDATE %s.auth
		SET
			email_verification_token=NULL
			,verified=@verified
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE
			email=@email
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"email":     email,
		"verified":  true,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *UserAuthStore) UpdateVerifiedToken(ctx context.Context, email domain.Email, token string) error {
	query := fmt.Sprintf(`
		UPDATE %s.auth
		SET
			email_verification_token=@token
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE
			email=@email
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"token":     token,
		"email":     email,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.ErrDataNotFound
		}

		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *UserAuthStore) UpdatePassword(ctx context.Context, email domain.Email, password domain.Password) error {
	query := fmt.Sprintf(`
		UPDATE %s.auth
		SET
			password_reset_token=NULL
			,password=@password
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE email = @email 
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"password": password,
		"email":    email,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.ErrDataNotFound
		}

		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func (pg *UserAuthStore) UpdateResetPasswordToken(ctx context.Context, email domain.Email, token string) error {
	query := fmt.Sprintf(`
		UPDATE %s.auth
		SET
			password_reset_token=@token
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE
			email=@email
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"token":     token,
		"email":     email,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.ErrDataNotFound
		}

		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *UserAuthStore) Select(ctx context.Context, email domain.Email) (*domain.UserAuth, error) {
	query := fmt.Sprintf(`
		SELECT 
			auth.auth_id
			,auth.verified
			,auth.email 
			,auth.oauth
			,auth.password
		FROM %s.auth
		WHERE email = @email;
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"email": email,
	}

	var user = &domain.UserAuth{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Verified,
		&user.Email,
		&user.OAuth,
		&user.Password,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}

		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	return user, nil
}
