package auth

import (
	"context"
	"fmt"

	"github.com/server/internal/core/domain/auth"
	"github.com/server/internal/core/domain/entity"
	"github.com/server/internal/core/domain/user"
)

type GetAccountInput struct {
	Email string
}

type GetAccountOutput struct {
	Username  user.Username
	Email     auth.Email
	Id        entity.Id
	CreatedBy entity.CreatedBy
	CreatedDT entity.CreatedDT
	UpdatedBy entity.UpdatedBy
	UpdatedDT entity.UpdatedDT
}

func (s *Service) GetAccount(ctx context.Context, account GetAccountInput) (*GetAccountOutput, error) {

	email, err := auth.NewEmail(account.Email)

	if err != nil {
		return nil, fmt.Errorf("invalid email supplied: %w", err)
	}

	user, err := s.userStore.GetAuthUser(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return &GetAccountOutput{
		Username:  user.Username,
		Email:     user.Email,
		Id:        user.Id,
		CreatedBy: user.CreatedBy,
		CreatedDT: user.CreatedDT,
		UpdatedBy: user.UpdatedBy,
		UpdatedDT: user.UpdatedDT,
	}, nil
}
