package userAuth

import (
	"github.com/server/internal/core/domain/auth"
	"github.com/server/internal/core/domain/entity"
	"github.com/server/internal/core/domain/user"
)

type UserAuth struct {
	Username user.Username
	Password auth.Password
	Email    auth.Email
}

type UserAuthBase struct {
	Username user.Username
	Password auth.Password
	Email    auth.Email
	Id       entity.Id
}

type UserAuthEntity struct {
	Username  user.Username
	Password  auth.Password
	Email     auth.Email
	Id        entity.Id
	CreatedBy entity.CreatedBy
	CreatedDT entity.CreatedDT
	UpdatedBy entity.UpdatedBy
	UpdatedDT entity.UpdatedDT
}

func NewBase(username user.Username, email auth.Email, password auth.Password, id entity.Id) UserAuthBase {
	return UserAuthBase{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}
}

func New(username user.Username, email auth.Email, password auth.Password) UserAuth {
	return UserAuth{
		Username: username,
		Email:    email,
		Password: password,
	}
}
