package auth

import (
	"context"
	"errors"

	"github.com/server/internal/core/ports"
)

var errIncorrectPassword = errors.New("password is incorrect")

type AuthService interface {
	LoginAccount(ctx context.Context, account LoginAccountInput) (*LoginAccountOuput, error)
	CreateAccount(ctx context.Context, account CreateAccountInput) (*CreateAccountOutput, error)
	UpdateAccount(ctx context.Context, account UpdateAccountInput) (*UpdateAccountOutput, error)
	RemoveAccount(ctx context.Context, account RemoveAccountInput) (*RemoveAccountOutput, error)
	GetAccount(ctx context.Context, account GetAccountInput) (*GetAccountOutput, error)
}

type Service struct {
	userStore ports.UserStore
}

func NewService(userStore ports.UserStore) *Service {
	return &Service{
		userStore: userStore,
	}
}
