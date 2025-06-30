package namespace

import "errors"

var (
	ErrNameSpaceInvalidID     = errors.New("namespace id is invalid")
	ErrNameSpaceNotExist      = errors.New("namespace doesn't exist")
	ErrNameSpaceConflict      = errors.New("namespace name already exist")
	ErrNameSpaceInvalidDetail = errors.New("invalid namespace detail")
)