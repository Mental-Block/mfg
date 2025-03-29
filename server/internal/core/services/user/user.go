package user

import (
	"context"

	"github.com/server/internal/core/ports"
)

type UserService interface {
	GetUser(ctx context.Context, user GetUserInput) (GetUserOuput, error)
	GetUsers(ctx context.Context) (GetUsersOuput, error)
}

type Service struct {
	userStore ports.UserStore
}

func NewService(userStore ports.UserStore) *Service {
	return &Service{
		userStore: userStore,
	}
}
