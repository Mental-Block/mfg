package auth

import (
	"context"
	"fmt"

	"github.com/server/internal/core/domain/auth"
	"github.com/server/internal/core/domain/entity"
	"github.com/server/internal/core/domain/user"
	"github.com/server/internal/core/domain/userAuth"
)

type UpdateAccountInput struct {
	Id          int
	Username    string
	Email       string
	OldPassword string
	Password    string
}

type UpdateAccountOutput = user.UserEntity

func (s *Service) UpdateAccount(ctx context.Context, account UpdateAccountInput) (*UpdateAccountOutput, error) {

	newUsername, err := user.NewUsername(account.Username)

	if err != nil {
		return nil, fmt.Errorf("invalid username supplied: %w", err)
	}

	newEmail, err := auth.NewEmail(account.Email)

	if err != nil {
		return nil, fmt.Errorf("invalid email supplied: %w", err)
	}

	authUser, err := s.userStore.GetAuthUser(ctx, newEmail)

	if err != nil {
		return nil, err
	}

	if isVerified := auth.VerifyPassword([]byte(account.OldPassword), string(authUser.Password)); !isVerified {
		return nil, errIncorrectPassword
	}

	newPassword, err := auth.NewPassword(account.Password)

	if err != nil {
		return nil, fmt.Errorf("invalid password supplied: %w", err)
	}

	user, err := s.userStore.UpdateUser(
		context.Background(),
		userAuth.NewBase(newUsername, newEmail, newPassword, entity.Id(account.Id)),
	)

	if err != nil {
		return nil, fmt.Errorf("failed update user: %w", err)
	}

	return user, nil
}
