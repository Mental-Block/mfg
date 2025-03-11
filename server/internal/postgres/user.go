package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id int
	Username string
	Email string
	Password string
	CreatedBy string
	CreatedDT string
	UpdatedBy string
	UpdatedDT string
}

func (pg *Repo) DeleteUser (ctx context.Context, user *User) error {
	query := fmt.Sprintf(`
		DELETE FROM %s.user
		WHERE id = @id;
	`, schema);

	args := pgx.NamedArgs {
		"id": user.Id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return fmt.Errorf("unable to delete row: %w", err)
	}

	return nil
}

func (pg *Repo) InsertUser(ctx context.Context, user *User) error {
	query := fmt.Sprintf(`
		INSERT INTO %s.user
		(
			username
			,password
			,email
		) 
		VALUES 
		(
			@username 
			,@password
			,@email
		);
	`, schema);

	args := pgx.NamedArgs {
		"username": user.Username,
		"email": user.Email,
		"password": user.Password,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Repo) UpdateUser (ctx context.Context, user *User) error {
	query := fmt.Sprintf(`
		UPDATE %s.user
		SET
			,username=@username
			,password=@password
			,email=@email
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE id = @id;
	`, schema);

	args := pgx.NamedArgs {
		"username": user.Username,
		"email": user.Email,
		"password": user.Password,
		"updatedBy": user.Username,
		"updatedDT": time.Now().UTC(),
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func (pg *Repo) GetUser(ctx context.Context, id int) (*User, error) {
	query := fmt.Sprintf(`
		SELECT 
			id
			,username
			,email 
			,created_by
			,created_dt
			,updated_by
			,updated_dt
		FROM %s.user 
		WHERE id = @id;
	`, schema);

	user := &User{}

	err := pg.db.QueryRow(ctx, query).Scan(user.Id, user.Username, user.Email, user.CreatedBy, user.CreatedDT, user.UpdatedBy, user.CreatedDT)

	if err != nil {
		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	return user, nil
}

func (pg *Repo) GetUsers(ctx context.Context, name string) ([]User, error) {
	query := fmt.Sprintf(`
		SELECT 
			id
			,username
			,email 
			,created_by
			,created_dt
			,updated_by
			,updated_dt
		FROM %s.user
		LIMIT 100;
	`,schema)

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[User])
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
