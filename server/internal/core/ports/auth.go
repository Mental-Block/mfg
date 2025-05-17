package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type AuthService interface {
	Permission(ctx context.Context, token string) (*string, *string, *domain.UserAuth, error)
	Register(ctx context.Context, email string, username string, password string) error
	RegisterFinish(ctx context.Context, token string) error
	Login(ctx context.Context, email string, password string) (*string, *string, *domain.UserAuth, error) 
	Verify(ctx context.Context, token string) (*string, error)
	UpdatePassword(ctx context.Context, token string, password string) error
	ResetPassword(ctx context.Context, email string) error
}

type AuthStore interface {
	UpdatePassword(ctx context.Context, email domain.Email, password domain.Password) error
	Select(ctx context.Context, email domain.Email) (*domain.Auth, error)
	SelectVersion(ctx context.Context, id domain.Id) (*int, error)
	// auth_id on the user table will need to be NULL before this func will work or you'll get foriegn key error, you can do this on the userStore side.
	Delete(ctx context.Context, id domain.Id) (*domain.Id, error)
	SelectCache(ctx context.Context, email domain.Email) (*domain.CachedUser, error)
	InsertCache(ctx context.Context, email domain.Email, password domain.Password, username domain.Username, token string) error
	DeleteCache(ctx context.Context, email domain.Email) error
}
