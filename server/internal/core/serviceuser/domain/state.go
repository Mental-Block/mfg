package domain

import (
	"time"
)

const (
	DefaultKeyType = "sv_rsa"
)

type State string

func (s State) String() string {
	return string(s)
}

const (
	Enabled  State = "enabled"
	Disabled State = "disabled"
)


type Secret struct {
	Id        string
	Title     string
	Value     string
	CreatedAt time.Time
}

type Token struct {
	Id       string
	Title     string
	Value     string
	CreatedAt time.Time
}
