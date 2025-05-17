package services

import (
	"context"
	"fmt"

	"github.com/server/internal"
	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)


type UserService struct {
	userStore ports.UserStore
}

func NewUserService(
	user ports.UserStore,
) *UserService {
	return &UserService{
		userStore: user,
	}
}

func (s *UserService) New(ctx context.Context, username string, email string, password string, oauth bool) (*domain.Id, error) {

	validUsername, err := domain.NewUsername(username)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid username supplied: %s", err.Error()))
	}

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid email supplied: %s", err.Error()) )
	}

	validPassword, err := domain.NewPassword(password)

	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid password supplied: %s", err.Error()))
	}

	id, err := s.userStore.Insert(ctx, validUsername, validEmail, validPassword, false)
    
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserStore.Error())
	}

	return id, nil
}

func (s *UserService) Remove(ctx context.Context, id int) (*domain.Id, error) {
	
	validId := domain.NewId(id)

	userId, err := s.userStore.Delete(ctx, validId)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserStore.Error())
	}

	return userId, nil
}

func (s *UserService) Get(ctx context.Context, id int) (*domain.User, error) {

	validId := domain.NewId(id)

	user, err := s.userStore.Select(ctx, validId)

	if err != nil {
		if err.Error() == domain.ErrUsersNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error())
		}
		
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserStore.Error())
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, id int, username string) (*domain.User, error) {
	
	validId := domain.NewId(id)
	
	validUsername, err := domain.NewUsername(username)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid username supplied: %s", err))
	}

	user, err := s.userStore.Update(ctx, validId, validUsername)

	if err != nil {
		if err.Error() == domain.ErrUsersNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error())
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("invalid username supplied: %s", err))
	}

	return user, nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	
	data, err := s.userStore.SelectUsers(ctx)

	if err != nil {
		if err.Error() == domain.ErrUsersNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error())
		}
	
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserStore.Error())
	}

	return data, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return nil,  internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid email supplied: %s", err))
	}

	user, err := s.userStore.SelectByEmail(ctx, validEmail)

	if err != nil {
		if err.Error() == domain.ErrUserNotFound.Error() {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrUserNotFound.Error())
		}
	
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserStore.Error())
	}

	return user, nil
}
