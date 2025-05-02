package domain

import (
	"errors"
	"net/mail"
	"strings"
)

type Email string

var (
	maxEmailLength = 255
)

var (
	ErrEmptyEmail         = errors.New("empty email supplied")
	ErrInvalidEmailFormat = errors.New("does not match email address format")
	ErrEmailToolong       = errors.New("email supplied is too long. Max 255 Chars")
)

func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return "", ErrEmptyEmail
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return "", ErrInvalidEmailFormat
	}

	if len(email) >= maxEmailLength {
		return "", ErrEmailToolong
	}

	return Email(email), nil
}
