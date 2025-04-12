package store

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type UserStore struct {
	db *postgres.Store
}

func NewUserStore(db *postgres.Store) *UserStore {
	return &UserStore{
		db,
	}
}

func (pg *UserStore) Select(ctx context.Context, email domain.Email) (*domain.User, error) {
	query := fmt.Sprintf(`
		SELECT 
			usr.user_id
			,usr.username
			,auth.verified
			,auth.email 
			,auth.oauth
		FROM %s.auth
			INNER JOIN %s.user AS usr ON usr.auth_id = auth.auth_id
		WHERE email = @email;
	`, postgres.Schema, postgres.Schema)

	args := pgx.NamedArgs{
		"email": email,
	}

	var user = &domain.User{}

	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Username,
		&user.Verified,
		&user.Email,
		&user.OAuth,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	return user, nil
}

func (pg *UserStore) Insert(ctx context.Context, email domain.Email, password domain.Password, username domain.Username, oauth bool) (*domain.Id, error) {
	authQuery := fmt.Sprintf(`
		INSERT INTO %v.auth
		(
			email
			,password
			,oauth
			,verified
			,created_by
			,created_dt
		)
		VALUES 
		(
			@email
			,@password
			,@oauth
			,@verified
			,@createdBy
			,@createdDT
		)
		RETURNING 
			auth_id;
	`, postgres.Schema)

	createdBy := domain.NewCreatedBy()
	createdDT := domain.NewCreatedDT()

	authArgs := pgx.NamedArgs{
		"password":  password,
		"email":     email,
		"oauth":     oauth,
		"verified":  oauth, // we can use oauth val to mark if verified
		"createdBy": createdBy,
		"createdDT": createdDT,
	}

	userquery := fmt.Sprintf(`
		INSERT INTO %v.user
		(
			auth_id
			,username
			,created_by
			,created_dt
		)
		VALUES
		(
			@authId
			,@username
			,@createdBy
			,@createdDT
		)
		RETURNING
			user_id;
	`, postgres.Schema)

	tx, err := pg.db.Pool.Begin(ctx)

	if err != nil {
		slog.Info(err.Error())
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	var authId *domain.Id
	err = tx.QueryRow(ctx, authQuery, authArgs).Scan(&authId)

	if err != nil {
		slog.Info(err.Error())
		return nil, fmt.Errorf("unable to insert row, rolling back transaction: %w", err)
	}

	userArgs := pgx.NamedArgs{
		"authId":    authId,
		"username":  username,
		"createdBy": createdBy,
		"createdDT": createdDT,
	}

	var id *domain.Id
	err = pg.db.QueryRow(ctx, userquery, userArgs).Scan(&id)

	if err != nil {
		slog.Info(err.Error())
		return nil, fmt.Errorf("unable to insert row, rolling back transaction: %w", err)
	}

	tx.Commit(ctx)

	return id, nil
}

func (pg *UserStore) Delete(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`
	DO $$
		DECLARE authId integer;
		BEGIN
			DELETE FROM %v.user 
			WHERE user_id = @id
			RETURNING auth_id into authId;
			
			DELETE FROM %v.auth 
			WHERE auth_id = authId;
	END $$;
	`, postgres.Schema, postgres.Schema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return nil, fmt.Errorf("unable to delete row: %w", err)
	}

	return &id, nil
}
