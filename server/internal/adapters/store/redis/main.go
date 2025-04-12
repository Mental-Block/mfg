package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	db *redis.Client
}

var (
	rdsInstance *Store
	rdsOnce     sync.Once
)

func NewStore(ctx context.Context, dbURL string) (*Store, error) {
	var err error

	rdsOnce.Do(func() {
		opts, dbErr := redis.ParseURL(dbURL)

		if dbErr != nil {
			err = fmt.Errorf("unable to create connection pool: %w", dbErr)
			return
		}

		db := redis.NewClient(opts)

		rdsInstance = &Store{db}
	})

	return rdsInstance, err
}

func (r *Store) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.db.Set(ctx, key, value, ttl).Err()
}

func (r *Store) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := r.db.Get(ctx, key).Result()
	bytes := []byte(res)
	return bytes, err
}

func (rds *Store) Delete(ctx context.Context, key string) error {
	return rds.db.Del(ctx, key).Err()
}

func (rds *Store) DeleteByPrefix(ctx context.Context, prefix string) error {
	var cursor uint64
	var keys []string

	for {
		var err error
		keys, cursor, err = rds.db.Scan(ctx, cursor, prefix, 100).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			err := rds.db.Del(ctx, key).Err()
			if err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

func (rds *Store) Ping(ctx context.Context) error {
	cmd := rds.db.Ping(ctx)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (rds *Store) Close() error {
	return rds.db.Close()
}
