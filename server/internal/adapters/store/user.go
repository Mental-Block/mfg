package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/server/internal/core/domain/auth"
	"github.com/server/internal/core/domain/entity"
	"github.com/server/internal/core/domain/user"
	"github.com/server/internal/core/domain/userAuth"
)

func (pg *Store) DeleteUser(ctx context.Context, id entity.Id) (*entity.Id, error) {
	query := fmt.Sprintf(`
		DELETE FROM %s.user
		WHERE id = @id;
	`, schema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		return nil, fmt.Errorf("unable to delete row: %w", err)
	}

	return &id, nil
}

func (pg *Store) InsertUser(ctx context.Context, userAuth userAuth.UserAuth) (*user.UserEntity, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s.user
		(
			username
			,password
			,email
			,created_dt
			,created_by
		) 
		VALUES 
		(
			@username 
			,@password
			,@email
			,@createdDT
			,@createdBy
		)
		RETURNING 
			id
			,username
			,created_by
			,created_dt
			,updated_by
			,updated_dt;
	`, schema)

	args := pgx.NamedArgs{
		"username":  userAuth.Username,
		"email":     userAuth.Email,
		"password":  userAuth.Password,
		"createdDT": entity.NewCreatedDT(),
		"createdBy": entity.NewCreatedBy(string(userAuth.Username)),
	}

	createdUser := &user.UserEntity{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&createdUser.Id,
		&createdUser.Username,
		&createdUser.CreatedBy,
		&createdUser.CreatedDT,
		&createdUser.UpdatedBy,
		&createdUser.UpdatedDT,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to insert row: %w", err)
	}

	return createdUser, nil
}

func (pg *Store) UpdateUser(ctx context.Context, userAuth userAuth.UserAuthBase) (*user.UserEntity, error) {
	query := fmt.Sprintf(`
		UPDATE %s.user
		SET
			username=@username
			,password=@password
			,email=@email
			,updated_by=@updatedBy
			,updated_dt=@updatedDT
		WHERE id = @id
		RETURNING 
			id
			,username
			,created_by
			,created_dt
			,updated_by
			,updated_dt;
	`, schema)

	args := pgx.NamedArgs{
		"id":        userAuth.Id,
		"username":  userAuth.Username,
		"email":     userAuth.Email,
		"password":  userAuth.Password,
		"updatedBy": entity.NewUpdatedBy(string(userAuth.Username)),
		"updatedDT": entity.NewUpdatedDT(),
	}

	updatedUser := &user.UserEntity{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&updatedUser.Id,
		&updatedUser.Username,
		&updatedUser.CreatedBy,
		&updatedUser.CreatedDT,
		&updatedUser.UpdatedBy,
		&updatedUser.UpdatedDT,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to update row: %w", err)
	}

	return updatedUser, nil
}

func (pg *Store) GetUser(ctx context.Context, id entity.Id) (*user.UserEntity, error) {
	query := fmt.Sprintf(`
		SELECT 
			id
			,username
			,created_by
			,created_dt
			,updated_by
			,updated_dt
		FROM %s.user 
		WHERE id = @id;
	`, schema)

	args := pgx.NamedArgs{
		"id": id,
	}

	var userEnt = &user.UserEntity{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&userEnt.Id,
		&userEnt.Username,
		&userEnt.CreatedBy,
		&userEnt.CreatedDT,
		&userEnt.UpdatedBy,
		&userEnt.UpdatedDT,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	return userEnt, nil
}

func (pg *Store) GetUsers(ctx context.Context) ([]user.UserEntity, error) {
	query := fmt.Sprintf(`
		SELECT 
			id
			,username
			,created_by
			,created_dt
			,updated_by
			,updated_dt
		FROM %s.user
		LIMIT 100;
	`, schema)

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[user.UserEntity])
}

// This will return user's email and password from database be careful with its useage
func (pg *Store) GetAuthUser(ctx context.Context, email auth.Email) (*userAuth.UserAuthEntity, error) {
	query := fmt.Sprintf(`
		SELECT 
			id
			,username
			,password
			,email 
			,created_by
			,created_dt
			,updated_by
			,updated_dt
		FROM %s.user 
		WHERE email = @email;
	`, schema)

	args := pgx.NamedArgs{
		"email": email,
	}

	var userAuth = &userAuth.UserAuthEntity{}

	err := pg.db.QueryRow(ctx, query, args).Scan(
		&userAuth.Id,
		&userAuth.Username,
		&userAuth.Password,
		&userAuth.Email,
		&userAuth.CreatedBy,
		&userAuth.CreatedDT,
		&userAuth.UpdatedBy,
		&userAuth.UpdatedDT,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	return userAuth, nil
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
