package services

import "errors"

var (
	ErrEmailAlreadyInUse = errors.New("email address is already registered")
	ErrInvalidEmail      = errors.New("invalid email address")
	ErrResourceStore   	 = errors.New("resource store error")
	ErrPermissionStore   = errors.New("permission store error")
	ErrRoleStore         = errors.New("role store error")
	ErrUserStore         = errors.New("user store error")
	ErrUserRoleStore     = errors.New("user role store error")
	ErrAuthStore     	 = errors.New("auth store error")
 	ErrTokenStore        = errors.New("token store error")
	ErrEmailService      = errors.New("email service error")
	ErrSigningToken      = errors.New("problem occured when signing token")
	ErrInvalidToken      = errors.New("token not valid")
)

