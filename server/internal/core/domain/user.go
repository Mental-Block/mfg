package domain

import (
	"errors"
)

var ErrUserNotFound = errors.New("user does not exist")
var ErrUsersNotFound = errors.New("users does not exist")

type User struct {
	Id Id
	Username Username
}

type UserAuth struct {
	Id 	Id
	Username Username
	Roles 	Roles
}
