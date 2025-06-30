package domain

import (
	"fmt"
	"strings"

	"github.com/server/pkg/utils"
)

var (
	ErrInvalidUsernameFormat = "invalid username supplied"
	maxUsernameLength = 30
)

type Username string

func (u Username) String() string {
	return string(u)
}

func NewUsername(u string) (Username, error) {
	u = strings.TrimSpace(u)

	if u == "" {
		return "", fmt.Errorf("empty username field")
	}

	if !utils.IsAlphanumeric(u) {
		return "", fmt.Errorf("expects only letters and numbers")
	}

	if !(len(u) <= maxUsernameLength) {
		return "", fmt.Errorf("title is max %v characters", maxUsernameLength)
	}

	return Username(u), nil
}