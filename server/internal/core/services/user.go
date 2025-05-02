package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)

type UserService struct {
	userProfileStore ports.UserProfileStore
	userStore        ports.UserStore
}

func NewUserService(
	user ports.UserStore,
	profile ports.UserProfileStore,
) *UserService {
	return &UserService{
		userProfileStore: profile,
		userStore:        user,
	}
}

func (s *UserService) GetProfile(ctx context.Context, id int) (*domain.UserProfile, error) {

	validId := domain.NewId(id)

	user, err := s.userProfileStore.Select(ctx, validId)

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrUserProfileStore
	}

	return user, nil
}

func (s *UserService) GetProfiles(ctx context.Context) ([]domain.UserProfile, error) {
	data, err := s.userProfileStore.SelectUsers(ctx)

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrUserProfileStore
	}

	return data, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, id int, username string) (*domain.UserProfile, error) {
	validId := domain.NewId(id)
	validUsername, err := domain.NewUsername(username)

	if err != nil {
		return nil, fmt.Errorf("invalid username supplied: %w", err)
	}

	user, err := s.userProfileStore.Update(ctx, validId, validUsername)

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrUserProfileStore
	}

	return user, nil
}

func (s *UserService) RemoveUser(ctx context.Context, id int) (*domain.Id, error) {
	validId := domain.NewId(id)

	userId, err := s.userStore.Delete(ctx, validId)

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrUserStore
	}

	return userId, nil
}

func (s *UserService) GetUser(ctx context.Context, email string) (*domain.User, error) {

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return nil, err
	}

	user, err := s.userStore.Select(ctx, validEmail)

	if err == domain.ErrDataNotFound {
		return nil, nil
	}

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrUserStore
	}

	return user, nil
}
