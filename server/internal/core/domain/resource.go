package domain

import (
	"errors"
	"strings"
)

var ErrResourceNotFound = errors.New("resource does not exist")

type Resource struct {
	Id Id
	Name string
//	Attribute map[string]any
}

var (
	ErrEmptyResource                	= errors.New("empty resource")
	ErrResourceTooLong               	= errors.New("max 50 characters")
	ErrResourceOnlyLetters 				= errors.New("expects only letters")
)

var maxResourceLength = 50

func NewResource(name string) (string, error) {
	un := strings.TrimSpace(name)

	if un == "" {
		return "", ErrEmptyResource
	}

	if !isLetter(un) {
		return "", ErrResourceOnlyLetters
	}

	if !(len(un) <= maxResourceLength) {
		return "", ErrResourceTooLong
	}

	return un, nil
}