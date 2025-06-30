package domain

import (
	"fmt"
	"strings"
)


var (
	ErrInvalidAvtarFormat = "invalid avtar supplied"
	maxAvtarLength = 256
)

type Avtar string

func (a Avtar) String() string {
	return string(a)
}

func NewAvtar(avtar string) (Avtar, error) {
	avtar = strings.TrimSpace(avtar)

	if avtar == "" {
		return "",  fmt.Errorf("empty avtar field")
	}

	if !(len(avtar) <= maxAvtarLength) {
		return "", fmt.Errorf("title is max %v characters", maxAvtarLength)
	}

	return Avtar(avtar), nil
}