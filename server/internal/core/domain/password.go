package domain

import (
	"errors"
	"unicode"
)

type Password string

var (
	ErrIncorrectPassword         = errors.New("incorrect email or password")
	ErrPaswordTooShort           = errors.New("min 8 characters")
	ErrPaswordTooLong            = errors.New("max 64 characters")
	ErrPaswordMustContainNumber  = errors.New("must contain a number")
	ErrPaswordMustContainUpper   = errors.New("must contain a uppercase characters")
	ErrPaswordMustContainLower   = errors.New("must contain a lowercase character")
	ErrPaswordMustContainSpecial = errors.New("must contain a special character")
)

var (
	minPasswordLength = 8
	maxPasswordLength = 64
)

func NewPassword(password string) (Password, error) {
	nonasciiPas := []rune(password)

	if !(len(nonasciiPas) >= minPasswordLength) {
		return "", ErrPaswordTooShort
	}

	if !(len(nonasciiPas) <= maxPasswordLength) {
		return "", ErrPaswordTooLong
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range nonasciiPas {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return "", ErrPaswordMustContainUpper
	}

	if !hasLower {
		return "", ErrPaswordMustContainLower
	}

	if !hasNumber {
		return "", ErrPaswordMustContainNumber
	}

	if !hasSpecial {
		return "", ErrPaswordMustContainSpecial
	}

	return Password(nonasciiPas), nil
}
