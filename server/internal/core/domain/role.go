package domain

import (
	"errors"
	"strings"
	"unicode"
)

var ErrRoleNotFound = errors.New("role does not exist")

type Role struct {
	Id Id
	Name string
}

type Roles []string

var (
	ErrEmptyRole                 = errors.New("empty role")
	ErrRoleTooLong               = errors.New("max 50 characters")
	ErrRoleOnlyLetters 			 = errors.New("expects only letters")
)

var maxRoleLength = 50

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func NewRole(name string) (string, error) {
	un := strings.TrimSpace(name)

	if un == "" {
		return "", ErrEmptyRole
	}

	if !isLetter(un) {
		return "", ErrRoleOnlyLetters
	}

	if !(len(un) <= maxRoleLength) {
		return "", ErrRoleTooLong
	}

	return un, nil
}