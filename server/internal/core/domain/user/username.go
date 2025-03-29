package user

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrEmptyUsername                 = errors.New("empty username supplied")
	ErrUsernameTooLong               = errors.New("username supplied is too long. Max 30 Chars")
	ErrContainSpecialCharsInUsername = errors.New("username supplied expects only letters and numbers")
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

	// Validate it is a valid username.
	return Username(un), nil
}
