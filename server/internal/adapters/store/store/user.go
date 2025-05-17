package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
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

func (pg *UserStore) Select(ctx context.Context, id domain.Id) (*domain.User, error) {

	query := fmt.Sprintf(`
		SELECT 
			user_id
			,username
		FROM %s.user 
		WHERE user_id = @id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var user = &domain.User{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Username,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}

	return user, nil
}

func (pg *UserStore) Insert(ctx context.Context, username domain.Username, email domain.Email, password domain.Password, oauth bool) (*domain.Id, error) {
	var (
		id 		  *domain.Id
		err        error
	)

	conn, err := pg.db.Acquire(context.Background())

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
	}

	defer conn.Release()

	err = postgres.Transaction(ctx, conn.Conn(), func(tx pgx.Tx) error {
		aq := authQueries{ conn: tx }

		authId, err := aq.insertQ(ctx, email, password, oauth)

		if (err != nil) {
			return internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
		}
		
		rq := userQueries{conn: tx}
		
		userId, err := rq.insertQ(ctx, *authId, username)

		if err != nil {
			return internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
		}

		id = userId

		return nil
	})

	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error()) 
	}

	return id, nil
}

func (pg *UserStore) Delete(ctx context.Context, id domain.Id) (*domain.Id, error) {
	conn, err := pg.db.Acquire(context.Background())

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
	}

	defer conn.Release()

	err = postgres.Transaction(ctx, conn.Conn(), func(tx pgx.Tx) error {	
		rq := userQueries{conn: tx}
		
		authId, err := rq.deleteQ(ctx, id)
		
		if (err != nil) {
			return internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
		}
		
		// can authId be nil as value can be "unassigned"
		if (authId != nil) {
			aq := authQueries{ conn: tx }

			err = aq.deleteQ(ctx, *authId)	
			
			if err != nil {
				return internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
			}
		}

		return nil
	})

	if err != nil {
		if err.Error() == domain.ErrUserNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error()) 
		}

		if err.Error() == domain.ErrAuthNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrAuthNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, err.Error()) 
	}

	return &id, nil
}

/* TODO - update user to update active column */
func (pg *UserStore) Update(ctx context.Context, id domain.Id, username domain.Username) (*domain.User, error) {
	query := fmt.Sprintf(`
		UPDATE %s.user
		SET
			username=@username
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE user_id = @id
		RETURNING 
			user_id
			,username
	`, postgres.PublicSchema)
	
	args := pgx.NamedArgs{
		"id":        id,
		"username":  username,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	user := &domain.User{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Username,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to update row: %s", err.Error())) 
	}


	return user, nil
}

/* TODO - update attributes for authorization */

func (pg *UserStore) SelectByEmail(ctx context.Context, email domain.Email) (*domain.User, error) {
	query := fmt.Sprintf(`
		SELECT 
			usr.user_id
			,usr.username
		FROM %s.auth
			INNER JOIN %s.user AS usr ON usr.auth_id = auth.auth_id
		WHERE auth.email = @email;
	`, postgres.PublicSchema, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"email": email,
	}

	var user = &domain.User{}

	err := pg.db.Pool.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Username,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to get row: %s", err.Error())) 
	}


	return user, nil
}

func (pg *UserStore) SelectUsers(ctx context.Context) ([]domain.User, error) {
	query := fmt.Sprintf(`
		SELECT 
			user_id AS id
			,username
		FROM %s.user
		LIMIT 100;
	`, postgres.PublicSchema)

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUsersNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to query rows: %s", err.Error())) 
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.User])
}

type userQueries struct {
	conn postgres.StoreTX
} 

func (uq *userQueries) deleteQ(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`DELETE FROM %v.user WHERE user_id = @id RETURNING auth_id;`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var authId *domain.Id
	err := uq.conn.QueryRow(ctx, query, args).Scan(
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

func (uq *userQueries) insertQ(ctx context.Context, authId domain.Id, username domain.Username) (*domain.Id, error) {
	query := fmt.Sprintf(`
		INSERT INTO %v.user
		(
			auth_id
			,username
		)
		VALUES
		(
			@authId
			,@username
		)
		RETURNING
			user_id;
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"authId":    authId,
		"username":  username,
	}

	var id *domain.Id
	err := uq.conn.QueryRow(ctx, query, args).Scan(
		&id,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}


	return id, nil
}
