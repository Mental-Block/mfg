package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/adapters/store/redis"
	"github.com/server/internal/core/domain"
)

type AuthStore struct {
	db *postgres.Store
	cache *redis.Store
}

func NewAuthStore(db *postgres.Store, cache *redis.Store) *AuthStore {
	return &AuthStore{
		db,
		cache,
	}
}

func (s *AuthStore) UpdatePassword(ctx context.Context, email domain.Email, password domain.Password) error {
	query := fmt.Sprintf(`
		UPDATE %s.auth
		SET
			password=@password
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE email = @email 
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"password": password,
		"email":    email,
	}

	_, err := s.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.ErrDataNotFound
		}

		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func (s *AuthStore) SelectUser(ctx context.Context, email domain.Email) (*domain.AuthUser, error) {

	bytes, err := s.cache.Get(ctx, string(email));

	if (err != nil) {
		return nil, err
	}

	var user *domain.AuthUser
	err = json.Unmarshal(bytes, &user)
	
	if (err != nil) {
		return nil, err
	}

	return user, nil
}

func (s *AuthStore) InsertUser(ctx context.Context, email domain.Email, password domain.Password, username domain.Username, verifiedToken string) error {
	
	user := domain.AuthUser{
		Email: email,
		Password: password,
		Username: username,
		Token: verifiedToken,
	} 
	
	userBytes, err := json.Marshal(user)

	if (err != nil) {
		return err
	}

	err = s.cache.Set(ctx, string(email), userBytes, domain.EmailVerificationToken)

	if (err != nil) {
		return err
	}

	return nil
}

func (s *AuthStore) RemoveUser(ctx context.Context, email domain.Email) error {
	 err := s.cache.Delete(ctx, string(email))

	if (err != nil) {
		return err
	}

	return nil
}

// func (pg *UserAuthStore) UpdateVerifiedToken(ctx context.Context, email domain.Email, token string) error {
// 	query := fmt.Sprintf(`
// 		UPDATE %s.auth
// 		SET
// 			email_verification_token=@token
// 			,updated_by=@updatedBy
// 			,updated_dt=@updatedDT
// 		WHERE
// 			email=@email
// 	`, postgres.Schema)

// 	args := pgx.NamedArgs{
// 		"token":     token,
// 		"email":     email,
// 		"updatedBy": domain.NewUpdatedBy(),
// 		"updatedDT": domain.NewUpdatedDT(),
// 	}

// 	_, err := pg.db.Exec(ctx, query, args)

// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return domain.ErrDataNotFound
// 		}

// 		return fmt.Errorf("unable to insert row: %w", err)
// 	}

// 	return nil
// }

// func (pg *UserAuthStore) UpdateResetPasswordToken(ctx context.Context, email domain.Email, token string) error {
// 	query := fmt.Sprintf(`
// 		UPDATE %s.auth
// 		SET
// 			password_reset_token=@token
// 			,updated_by=@updatedBy
// 			,updated_dt=@updatedDT
// 		WHERE
// 			email=@email
// 	`, postgres.Schema)

// 	args := pgx.NamedArgs{
// 		"token":     token,
// 		"email":     email,
// 		"updatedBy": domain.NewUpdatedBy(),
// 		"updatedDT": domain.NewUpdatedDT(),
// 	}

// 	_, err := pg.db.Exec(ctx, query, args)

// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return domain.ErrDataNotFound
// 		}

// 		return fmt.Errorf("unable to insert row: %w", err)
// 	}

// 	return nil
// }

// func (pg *UserAuthStore) Select(ctx context.Context, email domain.Email) (*domain.User, error) {
// 	query := fmt.Sprintf(`
// 		SELECT 
// 			auth.auth_id
// 			,auth.email 
// 			,auth.password
// 		FROM %s.auth
// 		WHERE email = @email;
// 	`, postgres.Schema)

// 	args := pgx.NamedArgs{
// 		"email": email,
// 	}

// 	var user = &domain.User{}

// 	err := pg.db.QueryRow(ctx, query, args).Scan(
// 		&user.Id,
// 		&user.Email,
// 		&user.Password,
// 	)

// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return nil, domain.ErrDataNotFound
// 		}

// 		return nil, fmt.Errorf("unable to get user: %w", err)
// 	}

// 	return user, nil
// }
