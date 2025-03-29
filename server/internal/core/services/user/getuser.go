package user

import (
	"context"
	"fmt"

	"github.com/server/internal/core/domain/entity"
	"github.com/server/internal/core/domain/user"
)

type GetUserInput int

type GetUserOuput *user.UserEntity

func (s *Service) GetUser(ctx context.Context, id GetUserInput) (GetUserOuput, error) {

	user, err := s.userStore.GetUser(ctx, entity.Id(id))

	if err != nil {

		return nil, fmt.Errorf("user store error")
	}

	return user, nil
}
