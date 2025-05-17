package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
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

func (pg *UserAuthStore) Update(ctx context.Context, id domain.Id, authId domain.Id) error {
	query := fmt.Sprintf(`
		UPDATE %v.user
		SET    
			auth_id = @authId
		    updated_by = @updatedAt
		    updated_dt = @updatedBy
		FROM  
		WHERE  
			user_id = @id
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id":        id,
		"authId":    authId,
		"updatedAt": domain.NewUpdatedDT(),
		"updatedBy": domain.NewUpdatedBy(),
	}
	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error())
		}

		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to update row: %s", err.Error()))
	}

	return nil
}

func (pg *UserAuthStore) Remove(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`
		UPDATE %v.user AS new_usr
		SET    
			auth_id = NULL
		    updated_by = @updatedAt
		    updated_dt = @updatedBy
		FROM  (SELECT auth_id FROM %v.user WHERE user_id = @id FOR UPDATE) AS old_usr 
		WHERE  
			new_usr.auth_id = old_usr.auth_id
		RETURNING 
			old_usr.auth_id AS id
	`, postgres.PublicSchema, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id":        id,
		"updatedAt": domain.NewUpdatedDT(),
		"updatedBy": domain.NewUpdatedBy(),
	}

	var authId *domain.Id
	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&authId,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error()))
	}

	return authId, nil
}