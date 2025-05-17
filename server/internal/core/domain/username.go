package domain

import (
	"errors"
	"strings"
	"unicode"
)

type Username string

var (
	ErrEmptyUsername                 = errors.New("empty username")
	ErrUsernameTooLong               = errors.New("max 30 characters")
	ErrContainSpecialCharsInUsername = errors.New("expects only letters and numbers")
)

var (
	maxUsernameLength = 30
)

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func NewUsername(un string) (Username, error) {
	un = strings.TrimSpace(un)

	if un == "" {
		return "", ErrEmptyUsername
	}

	if !isAlphanumeric(un) {
		return "", ErrContainSpecialCharsInUsername
	}

	if !(len(un) <= maxUsernameLength) {
		return "", ErrUsernameTooLong
	}

	return Username(un), nil
}
