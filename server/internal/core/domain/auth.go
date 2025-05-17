package domain

import (
	"errors"
	"time"
)

var (
	ErrAuthNotFound = errors.New("no auth data found")
    ErrTokenNotFound = errors.New("no token data found")
)

var (
	RefreshTokenName = "mfg-refresh-token"
	AuthTokenName = "mfg-authourization"
)

var (
	PasswordResetTokenDuration = time.Minute * 5
	RefreshTokenDuration = time.Hour * 24 * 30
	AuthTokenDuration = time.Minute * 15
	EmailVerificationToken = time.Minute * 15
)

type Auth struct {
	Id 		 Id
	Email    Email
	Password Password
	OAuth	 bool
	Version  int
}

type CachedUser struct {
	Username Username
	Email    Email
	Password Password
	Token 	 string
}

