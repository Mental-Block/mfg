package services

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFailure = errors.New("internal failure")
)

var (
	ErrIncorrectPassword = errors.New("incorrect username or password")
	ErrUserStore         = errors.New("user store error")
	ErrUserProfileStore  = errors.New("user profile store error")
	ErrUserAuthStore     = errors.New("user auth store error")
 	ErrTokenStore        = errors.New("token store error")
	ErrEmailService      = errors.New("email service error")
	ErrSigningToken      = errors.New("problem occured when signing token")
	ErrInvalidToken      = errors.New("token not valid")
)
