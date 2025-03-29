package ports

import (
	"context"
	"errors"

	"github.com/server/internal/core/domain/auth"
	"github.com/server/internal/core/domain/entity"
	"github.com/server/internal/core/domain/user"
	"github.com/server/internal/core/domain/userAuth"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
)

type UserStore interface {
	DeleteUser(ctx context.Context, id entity.Id) (*entity.Id, error)
	GetUser(ctx context.Context, id entity.Id) (*user.UserEntity, error)
	GetUsers(ctx context.Context) ([]user.UserEntity, error)
	InsertUser(ctx context.Context, user userAuth.UserAuth) (*user.UserEntity, error)
	UpdateUser(ctx context.Context, user userAuth.UserAuthBase) (*user.UserEntity, error)
	GetAuthUser(ctx context.Context, email auth.Email) (*userAuth.UserAuthEntity, error)
}
