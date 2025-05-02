package store

import (
	"context"
	"time"

	"github.com/server/internal/adapters/store/redis"
)

type TokenStore struct {
	cache *redis.Store
}

func NewTokenStore(cache *redis.Store) *TokenStore {
	return &TokenStore{
		cache: cache,
	}
}

func (s *TokenStore) Insert(ctx context.Context, key string, token string, duration time.Duration) error {

	err := s.cache.Set(ctx, key, []byte(token), duration) 

	if (err != nil) {
		return err
	}

	return nil
}

func (s *TokenStore) Select(ctx context.Context, key string) (*string, error) {

	result, err := s.cache.Get(ctx, key)

	if (err != nil) {
		return nil, err
	}

	res := string(result)

	return &res, nil
}

func (s *TokenStore) Remove(ctx context.Context, key string) (error) {
	 err := s.cache.Delete(ctx, key)

	if (err != nil) {
		return err
	}

	return nil
}
 