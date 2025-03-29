package auth

import (
	"errors"
	"net/mail"
	"strings"
)

type Email = string

var (
	maxEmailLength = 255
)

var (
	ErrEmptyEmail   = errors.New("empty email supplied")
	ErrInvalidEmail = errors.New("does not match email address format")
	ErrEmailToolong = errors.New("email supplied is too long. Max 255 Chars")
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return "", ErrEmptyEmail
	}

	if !valid(email) {
		return "", ErrInvalidEmail
	}

	if len(email) >= maxEmailLength {
		return "", ErrEmailToolong
	}

	return Email(email), nil
}
