package auth

import (
	"context"
	"fmt"

	"github.com/server/internal/core/domain/auth"
	"github.com/server/internal/core/domain/user"
)

type LoginAccountInput struct {
	Email    string
	Password string
}

type LoginAccountOuput = user.UserEntity

func (s *Service) LoginAccount(ctx context.Context, account LoginAccountInput) (*LoginAccountOuput, error) {

	email, err := auth.NewEmail(account.Email)

	if err != nil {
		return nil, fmt.Errorf("invalid email supplied: %w", err)
	}

	user, err := s.userStore.GetAuthUser(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	if isVerified := auth.VerifyPassword([]byte(account.Password), string(user.Password)); !isVerified {
		return nil, errIncorrectPassword
	}

	return &LoginAccountOuput{
		Username:  user.Username,
		Id:        user.Id,
		CreatedBy: user.CreatedBy,
		CreatedDT: user.CreatedDT,
		UpdatedBy: user.UpdatedBy,
		UpdatedDT: user.UpdatedDT,
	}, nil
}
