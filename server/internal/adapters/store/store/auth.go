package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	redisV9 "github.com/redis/go-redis/v9"
	"github.com/server/internal"
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

func (pg* AuthStore) Delete(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`DELETE FROM %v.auth WHERE auth_id = @id`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error()))  
	}

	return &id, nil
}

func (pg *AuthStore) SelectVersion(ctx context.Context, id domain.Id) (*int, error) {
	query := fmt.Sprintf(`
		SELECT 
			version
		FROM %s.auth
		WHERE auth_id = @id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var version *int

	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&version,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	return version, nil
}

func (pg *AuthStore) Select(ctx context.Context, email domain.Email) (*domain.Auth, error) {
	query := fmt.Sprintf(`
		SELECT 
			auth_id
			,email 
			,password
			,oauth
			,version
		FROM %s.auth
		WHERE email = @email;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"email": email,
	}

	var auth = &domain.Auth{}

	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&auth.Id,
		&auth.Email,
		&auth.Password,
		&auth.OAuth,
		&auth.Version,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	return auth, nil
}

func (pg *AuthStore) UpdatePassword(ctx context.Context, email domain.Email, password domain.Password) error {
	query := fmt.Sprintf(`
		UPDATE %s.auth
		SET
			version=(version + 1)
			,password=@password
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE email = @email 
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"password": password,
		"email":    email,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error()) 
		}

		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to update row: %s", err.Error())) 
	}

	return nil
}

func (rd *AuthStore) SelectCache(ctx context.Context, email domain.Email) (*domain.CachedUser, error) {

	bytes, err := rd.cache.Get(ctx, string(email));

	if (err != nil) {
		if (err == redisV9.Nil) {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error()) 
	}

	var user *domain.CachedUser
	err = json.Unmarshal(bytes, &user)
	
	if (err != nil) {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error()) 
	}

	return user, nil
}

func (rd *AuthStore) InsertCache(ctx context.Context, email domain.Email, password domain.Password, username domain.Username, token string) error {
	
	user := domain.CachedUser{
		Email: email,
		Password: password,
		Username: username,
		Token: token,
	} 
	
	userBytes, err := json.Marshal(user)

	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error()) 
	}

	err = rd.cache.Set(ctx, string(email), userBytes, domain.EmailVerificationToken)

	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error()) 
	}

	return nil
}

func (rd *AuthStore) DeleteCache(ctx context.Context, email domain.Email) error {
	 err := rd.cache.Delete(ctx, string(email))

	if (err != nil) {
		if (err == redisV9.Nil) {
			return internal.WrapErrorf(err, internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error())
		}

		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error()) 
	}

	return nil
}

type authQueries struct {
	conn postgres.StoreTX
} 

func (aq *authQueries) deleteQ(ctx context.Context, id domain.Id) error {
	query := fmt.Sprintf(`DELETE FROM %v.auth WHERE auth_id = @id;`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := aq.conn.Exec(ctx, query, args)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error())) 
	}

	return nil
}

func (aq *authQueries) insertQ(ctx context.Context, email domain.Email, password domain.Password, oauth bool) (*domain.Id, error) {
	query := fmt.Sprintf(`
		INSERT INTO %v.auth
		(
			email
			,password
			,oauth
		)
		VALUES 
		(
			@email
			,@password
			,@oauth
		)
		RETURNING 
			auth_id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"password":  password,
		"email":     email,
		"oauth":     oauth,
	}

	var id *domain.Id
	err := aq.conn.QueryRow(ctx, query, args).Scan(
		&id,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}

	return id, nil 
}