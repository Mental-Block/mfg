package domain

import (
	"errors"
	"strings"
)

var ErrPermissionNotFound = errors.New("permission does not exist")

type Permission struct {
	Id	Id
	Name string
}

var (
	ErrEmptyPermission                	= errors.New("empty permission")
	ErrPermissionTooLong               	= errors.New("max 50 characters")
	ErrPermissionOnlyLetters 			= errors.New("expects only letters")
)

var maxPermissionLength = 50

func NewPermission(name string) (string, error) {
	un := strings.TrimSpace(name)

	if un == "" {
		return "", ErrEmptyPermission
	}

	if !isLetter(un) {
		return "", ErrPermissionOnlyLetters
	}

	if !(len(un) <= maxPermissionLength) {
		return "", ErrPermissionTooLong
	}

	return un, nil
}