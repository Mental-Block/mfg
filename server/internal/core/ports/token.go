package ports

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	Parse(token string) (jwt.MapClaims, error)
	Create(claims jwt.MapClaims) (*string, error)
	Verify(token string) error
}

type TokenStore interface {
	Select(ctx context.Context, key string) (*string, error)
	Insert(ctx context.Context, key string, token string, duration time.Duration) error
	Remove(ctx context.Context, key string) error
}
