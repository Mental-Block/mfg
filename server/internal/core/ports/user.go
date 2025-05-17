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
	New(ctx context.Context, username string, email string, password string, oauth bool) (*domain.Id, error)
	Remove(ctx context.Context, id int) (*domain.Id, error)
	Update(ctx context.Context, id int, username string) (*domain.User, error)
	Get(ctx context.Context, id int) (*domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error) 
}

type UserStore interface {
	// get a user by id
	Select(ctx context.Context, id domain.Id) (*domain.User, error)
	// delete a user and authentication with this user
	Delete(ctx context.Context, id domain.Id) (*domain.Id, error)
	// insert a new user and authentication with this user
	Insert(ctx context.Context, username domain.Username, email domain.Email, password domain.Password, oauth bool) (*domain.Id, error)
	// update user 
	Update(ctx context.Context, id domain.Id, username domain.Username) (*domain.User, error) 
	// get a user by email
	SelectByEmail(ctx context.Context, email domain.Email) (*domain.User, error)
	// get users
	SelectUsers(ctx context.Context) ([]domain.User, error)
}


