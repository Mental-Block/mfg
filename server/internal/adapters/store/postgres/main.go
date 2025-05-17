package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/server/internal"
)

type Store struct {
	*pgxpool.Pool
}

var (
	pgInstance *Store
	pgOnce     sync.Once
)

const PublicSchema = "public"

func NewStore(ctx context.Context, dbURL string) (*Store, error) {
	var err error

	pgOnce.Do(func() {
		db, dbErr := pgxpool.New(ctx, dbURL)

		if dbErr != nil {
			err = fmt.Errorf("unable to create connection pool: %w", dbErr)
			return
		}

		pgInstance = &Store{db}
	})

	return pgInstance, err
}

func (pg *Store) Ping(ctx context.Context) error {
	return pg.Pool.Ping(ctx)
}

func (pg *Store) Close() {
	pg.Pool.Close()
}

type StoreTX interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
}

// helper func for database transactions
func Transaction(ctx context.Context, conn *pgx.Conn, fn func(tx pgx.Tx) error) error {
	tx, err := conn.Begin(ctx)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeUnknown, err.Error())
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)

		return internal.NewErrorf(internal.ErrorCodeUnknown, err.Error())
	}

	if err := tx.Commit(ctx); err != nil {
		return internal.NewErrorf(internal.ErrorCodeUnknown, err.Error())
	}

	return nil
}