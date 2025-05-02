package ports

import (
	"context"
	"errors"

	"github.com/server/internal/core/domain"
)

var (
	ErrCouldntAuthenticate = errors.New("account could authenticate")
)

type UserService interface {
	GetProfile(ctx context.Context, id int) (*domain.UserProfile, error)
	GetProfiles(ctx context.Context) ([]domain.UserProfile, error)
	UpdateProfile(ctx context.Context, id int, username string) (*domain.UserProfile, error)
	RemoveUser(ctx context.Context, id int) (*domain.Id, error)
	GetUser(ctx context.Context, email string) (*domain.User, error)
}

type UserStore interface {
	SelectEmail(ctx context.Context, id domain.Id) (*domain.Email, error)
	Delete(ctx context.Context, id domain.Id) (*domain.Id, error)
	Insert(ctx context.Context, email domain.Email, password domain.Password, username domain.Username, oauth bool) (*domain.Id, error)
	Select(ctx context.Context, email domain.Email) (*domain.User, error)
}

type UserProfileStore interface {
	SelectUsers(ctx context.Context) ([]domain.UserProfile, error)
	Select(ctx context.Context, id domain.Id) (*domain.UserProfile, error)
	Update(ctx context.Context, id domain.Id, username domain.Username) (*domain.UserProfile, error)
}
