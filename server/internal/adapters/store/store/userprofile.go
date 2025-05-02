package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type UserProfileStore struct {
	db *postgres.Store
}

func NewUserProfileStore(db *postgres.Store) *UserProfileStore {
	return &UserProfileStore{
		db,
	}
}

func (pg *UserProfileStore) Update(ctx context.Context, id domain.Id, username domain.Username) (*domain.UserProfile, error) {
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
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"id":        id,
		"username":  username,
		"updatedBy": domain.NewUpdatedBy(),
		"updatedDT": domain.NewUpdatedDT(),
	}

	user := &domain.UserProfile{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Username,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to update row: %w", err)
	}

	return user, nil
}

func (pg *UserProfileStore) Select(ctx context.Context, id domain.Id) (*domain.UserProfile, error) {
	query := fmt.Sprintf(`
		SELECT 
			user_id
			,username
		FROM %s.user 
		WHERE user_id = @id;
	`, postgres.Schema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var user = &domain.UserProfile{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&user.Id,
		&user.Username,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}

		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	return user, nil
}

func (pg *UserProfileStore) SelectUsers(ctx context.Context) ([]domain.UserProfile, error) {
	query := fmt.Sprintf(`
		SELECT 
			user_id AS id
			,username
		FROM %s.user
		LIMIT 100;
	`, postgres.Schema)

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}

		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.UserProfile])
}

// func (pg *Repo) BulkInsertUsers(ctx context.Context, users []User) error {
//   query := `INSERT INTO users (name, email) VALUES (@userName, @userEmail)`

//   batch := &pgx.Batch{}
//   for _, user := range users {

//     args := pgx.NamedArgs{
//       "userName": user.Username,
//       "userEmail": user.Email,
//     }

//     batch.Queue(query, args)
//   }

//   results := pg.db.SendBatch(ctx, batch)
//   defer results.Close()

//   for _, user := range users {
//     _, err := results.Exec()
//     if err != nil {
//       var pgErr *pgconn.PgError

//       if errors.As(err, &pgErr) && pgErr.Code == pgErr.ConstraintName {
//           log.Printf("user %s already exists", user.Name)
//           continue
//       }

//       return fmt.Errorf("unable to insert row: %w", err)
//     }
//   }

//   return results.Close()
// }

// func (pg *Repo) CopyInsertUsers(ctx context.Context, users []User) error {
//   entries := [][]any{}
//   columns := []string{"name", "email"}
//   tableName := "user"

//   for _, user := range users {
//     entries = append(entries, []any{user.Name, user.Email})
//   }

//   _, err := pg.db.CopyFrom(
//     ctx,
//     pgx.Identifier{tableName},
//     columns,
//     pgx.CopyFromRows(entries),
//   )

//   if err != nil {
//     return fmt.Errorf("error copying into %s table: %w", tableName, err)
//   }

//   return nil
// }
