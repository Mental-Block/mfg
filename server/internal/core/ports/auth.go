package ports

import (
	"context"

	"github.com/server/internal/core/domain"
)

type PasswordService interface {
	Hash(password domain.Password) domain.Password
	Verify(password domain.Password, hash domain.Password) bool
}

type AuthService interface {
	/* could probably break out token into seperate service */
	NewAuthToken(id domain.Id, roles []string) (*string, error)
    NewEmailVerificationToken(email domain.Email) (*string, error) 
    NewPasswordResetToken(email domain.Email) (*string, error) 
	NewRefreshToken(id domain.Id) (*string, error)

	RegisterOAuth(ctx context.Context, email string) // Not implimented
	Register(ctx context.Context, email string, username string, password string) error
	FinishRegister(ctx context.Context, token string) error
	Verify(ctx context.Context, token string) (string, error)
	UpdatePassword(ctx context.Context, token string, password string) error
	ResetPassword(ctx context.Context, email string) error
	LoginOAuth(ctx context.Context) //Not implimented
	Login(ctx context.Context, email string, password string) (*string, error) 
	Logout(ctx context.Context, token string) error
}

type AuthStore interface {
	UpdatePassword(ctx context.Context, email domain.Email, password domain.Password) error
	SelectUser(ctx context.Context, email domain.Email) (*domain.AuthUser, error)
	InsertUser(ctx context.Context, email domain.Email, password domain.Password, username domain.Username, verifiedToken string) error
	RemoveUser(ctx context.Context, email domain.Email) error
}
