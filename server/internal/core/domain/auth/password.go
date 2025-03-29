package auth

import (
	"encoding/base64"
	"errors"
	"unicode"

	"github.com/server/env"

	"golang.org/x/crypto/argon2"
)

type Password = string

var (
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

func hashPassword(password []byte, salt []byte) string {
	return base64.StdEncoding.EncodeToString(argon2.Key(password, salt, 1, 32*1024, 4, 32))
}

func VerifyPassword(password []byte, hash string) bool {
	cfg := env.Env()

	newHash := hashPassword(password, []byte(cfg.Web.Salt))
	return newHash == hash
}

func NewPassword(password string) (Password, error) {
	cfg := env.Env()

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

	return Password(hashPassword([]byte((string(nonasciiPas))), []byte(cfg.Web.Salt))), nil
}
