package store

import (
	"context"
	"time"

	redisV9 "github.com/redis/go-redis/v9"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/redis"
	"github.com/server/internal/core/domain"
)

type TokenStore struct {
	cache *redis.Store
}

func NewTokenStore(cache *redis.Store) *TokenStore {
	return &TokenStore{
		cache: cache,
	}
}

func (rd *TokenStore) Insert(ctx context.Context, key string, token string, duration time.Duration) error {

	err := rd.cache.Set(ctx, key, []byte(token), duration) 

	if (err != nil) {
		return internal.WrapErrorf(err,internal.ErrorCodeUnknown, err.Error()) 
	}

	return nil
}

func (rd *TokenStore) Select(ctx context.Context, key string) (*string, error) {

	result, err := rd.cache.Get(ctx, key)

	if (err != nil) {
		if (err == redisV9.Nil) {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrTokenNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error()) 
	}

	res := string(result)

	return &res, nil
}

func (rd *TokenStore) Delete(ctx context.Context, key string) (error) {

	 err := rd.cache.Delete(ctx, key)

	if (err != nil) {
		if (err == redisV9.Nil) {
			return internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrTokenNotFound.Error())
		}

		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, err.Error()) 
	}

	return nil
}
 