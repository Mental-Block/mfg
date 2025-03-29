package user

import (
	"context"
	"fmt"

	"github.com/server/internal/core/domain/user"
)

type GetUsersOuput = []user.UserEntity

func (s *Service) GetUsers(ctx context.Context) (GetUsersOuput, error) {

	users, err := s.userStore.GetUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return users, nil
}
