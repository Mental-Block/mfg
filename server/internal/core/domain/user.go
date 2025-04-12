package domain

import (
	"errors"
)

var ErrUserNotFound = errors.New("user does not exist")

type UserAuth struct {
	Id       Id
	Verified bool
	OAuth    bool
	Email    Email
	Password Password
}

type UserProfile struct {
	Id       Id
	Username Username
}

type User struct {
	Id       Id
	Username Username
	Verified bool
	OAuth    bool
	Email    Email
}
