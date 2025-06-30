package session

import (
	"context"
	"time"

	"github.com/server/internal/adapters/store/postgres/auth"
	"github.com/server/internal/adapters/store/redis"
	"github.com/server/internal/core/auth/domain"
)

type ISessionStore interface { 
	Select(ctx context.Context, key domain.SessionKey) (*redis.SessionModel, error) 
	Insert(ctx context.Context, key domain.SessionKey, value domain.Session, duration time.Duration) error
	Delete(ctx context.Context, key domain.SessionKey) error
	DeleteByPrefix(ctx context.Context, key domain.SessionKey) error
	GenerateKey(id, authId string) domain.SessionKey
}

type IAuthStore interface {
	Select(ctx context.Context, email domain.Email) (*auth.AuthModel, error)
}