package domain

import (
	"errors"
	"unicode"
)

type Password string

var (
	ErrIncorrectPassword         = errors.New("incorrect email or password")
	ErrPaswordTooShort           = errors.New("password supplied is too short. Min 8 chars")
	ErrPaswordTooLong            = errors.New("password supplied is too long. Max 64 chars")
	ErrPaswordMustContainNumber  = errors.New("password supplied must contain a number")
	ErrPaswordMustContainUpper   = errors.New("password supplied must contain a uppercase character")
	ErrPaswordMustContainLower   = errors.New("password supplied must contain a lowercase character")
	ErrPaswordMustContainSpecial = errors.New("password supplied must contain a special character")
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
