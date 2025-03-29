package auth

import (
	"context"
	"fmt"

	"github.com/server/internal/core/domain/entity"
)

type RemoveAccountInput struct {
	Id int
}

type RemoveAccountOutput struct {
	Id entity.Id
}

func (s *Service) RemoveAccount(ctx context.Context, account RemoveAccountInput) (*RemoveAccountOutput, error) {
	_, err := s.userStore.DeleteUser(ctx, entity.Id(account.Id))

	if err != nil {
		return nil, fmt.Errorf("failed to delete a user: %w", err)
	}

	return &RemoveAccountOutput{Id: entity.Id(account.Id)}, nil
}
