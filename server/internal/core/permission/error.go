package permission

import "errors"

var (
	ErrInvalidID     = errors.New("permission id is invalid")
	ErrNotExist      = errors.New("permission doesn't exist")
	ErrInvalidDetail = errors.New("invalid permission detail")
	ErrPermissionService = errors.New("permission service error")
)
