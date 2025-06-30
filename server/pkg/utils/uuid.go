package utils

import (
	"errors"

	"github.com/google/uuid"
)

type UUID string 

var ErrInvalidUUIDFormat = "invalid UUID supplied: %s"

func NewUUID() UUID {
	return UUID(uuid.NewString())
}

func (i UUID) String() string {
	return string(i)
}

func ConvertStringToUUID(id string) (UUID, error) {
	if IsNullUUID(id) {
		return  UUID(""), errors.New("failed string to UUID parse")
	}

	return UUID(id), nil
}

// IsValidUUID returns true if passed string in uuid format
// defined by `github.com/google/uuid`.Parse
// else return false
func IsValidUUID(key string) bool {
	_, err := uuid.Parse(key)
	return err == nil
}

// IsNullUUID returns true if passed string is a null uuid or is not a valid uuid
// defined by `github.com/google/uuid`.Parse and `github.com/google/uuid`.Nil respectively
// else return false
func IsNullUUID(key string) bool {
	k, err := uuid.Parse(key)
	return err != nil || k == uuid.Nil
}
