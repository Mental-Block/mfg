package domain

import (
	"errors"
)

var ErrUserNotFound = errors.New("user does not exist")


type User struct {
	Id       Id
	Username Username
	Email    Email
	Password Password
	OAuth    bool
}

type UserProfile struct {
	Id       Id
	Username Username
}
