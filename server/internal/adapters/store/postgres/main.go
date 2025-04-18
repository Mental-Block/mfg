package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*pgxpool.Pool
}

var (
	pgInstance *Store
	pgOnce     sync.Once
)

const Schema = "public"

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
