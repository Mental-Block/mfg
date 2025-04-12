package ports

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/server/internal/core/domain"
)

type PasswordService interface {
	HashPassword(password domain.Password) domain.Password
	VerifyPassword(password domain.Password, hash domain.Password) bool
}

type TokenService interface {
	ParseToken(token string) (jwt.MapClaims, error)
	CreateToken(claims jwt.MapClaims) (*string, error)
	VerifyToken(token string) error
}

type AuthService interface {
	Login(ctx context.Context, email string, password string, oauth bool) (*string, error)
	Register(ctx context.Context, email string, username string, password string, oauth bool) (*string, error)
	VerifyUser(ctx context.Context, token string) error
	UpdatePassword(ctx context.Context, token string, password string) error
	ResetPassword(ctx context.Context, email domain.Email) error
}
