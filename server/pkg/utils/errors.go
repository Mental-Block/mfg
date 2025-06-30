package utils

import (
	"errors"
	"fmt"
)

// Error represents an error that could be wrapping another error, it includes a code for determining what
// triggered the error.
type Error struct {
	orig error
	msg  string
	code ErrorCode
}

// Enum for error codes
type ErrorCode uint

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeNotAuthorized
	ErrorCodeInvalidArgument
)	

// appends previous error to this new error
func WrapErrorf(orig error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		code: code,
		orig: orig,
		msg:  fmt.Sprintf(format, a...),
	}
}

// creates a new error to bubble up
func NewErrorf(code ErrorCode, format string, a ...interface{}) error {
	return WrapErrorf(nil, code, format, a...)
}

func (e *Error) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.orig)
	}

	return e.msg
}

// returns the error at the top of the stack
func (e *Error) Top() error {
	return errors.New(e.msg)
}

// returns the orignial error in the stack
func (e *Error) Unwrap() error {
	return e.orig
}

// returns the error code
func (e *Error) Code() ErrorCode {
	return e.code
}
