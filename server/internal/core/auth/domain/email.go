package domain

import (
	"errors"
	"net"
	"net/mail"
	"strings"
)

type Email string

func (e Email) String() string{
	return string(e)
}

var (
	maxEmailLength = 255
)

var (
	ErrEmptyEmail         = errors.New("empty email supplied")
	ErrInvalidEmailFormat = errors.New("does not match email address format")
	ErrEmailToolong       = errors.New("email supplied is too long. Max 255 Chars")
	ErrInvalidEmail       = errors.New("invalid email address")
	ErrNoEmailFound       = errors.New("no email found")
)

func (e Email) NewEmail() (Email, error) {
	email := strings.ToLower(strings.TrimSpace(e.String()))

	if email == "" {
		return "", ErrEmptyEmail
	}

	if _, err := mail.ParseAddress(e.String()); err != nil {
		return "", ErrInvalidEmailFormat
	}

	if len(email) >= maxEmailLength {
		return "", ErrEmailToolong
	}

	return Email(email), nil
}

func (e Email) DNSLookUp() (error) {

	parts := strings.Split(e.String(), "@")

	if (len(parts) != 2) {
		return ErrInvalidEmail
	} 

	mx, err := net.LookupMX(parts[1])

	if (err != nil ) {
		return err
	}

  if (len(mx) == 0) {
    return ErrNoEmailFound
  }

  return nil
}