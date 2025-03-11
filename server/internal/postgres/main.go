package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
  db *pgxpool.Pool
}

type Results[T any] struct {
    Data []T
}

var (
  pgInstance *Repo
  pgOnce     sync.Once
)

const schema string = "public";

func NewRepo(ctx context.Context, dbURL string) (*Repo, error) {
  var err error

  pgOnce.Do(func() {
    db, dbErr := pgxpool.New(ctx, dbURL)
    
    if dbErr != nil {
      err = fmt.Errorf("unable to create connection pool: %w", dbErr)
      return
    }

    pgInstance = &Repo{db}
  })

  return pgInstance, err
}

func (pg *Repo) Ping(ctx context.Context) error {
  return pg.db.Ping(ctx)
}

func (pg *Repo) Close() {
  pg.db.Close()
}

