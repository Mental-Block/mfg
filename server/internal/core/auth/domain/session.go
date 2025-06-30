package domain

import (
	"time"

	"github.com/server/pkg/utils"
)

type SessionKey string

func (s SessionKey) String() string {
	return string(s)
}

type Session struct {
	Id utils.UUID
	AuthId utils.UUID
	Email Email
	Version int
	CreatedAt time.Time
	ExpiresAt time.Time
}
