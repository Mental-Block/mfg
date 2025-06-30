package domain

import (
	"fmt"
	"strings"

	"github.com/server/pkg/utils"
)


var (
	maxTitleLength = 50
	ErrInvalidTileFormat = "invalid title supplied"
)

type Title string

func (t Title) String() string {
	return string(t)
}

func NewTitle(t string) (Title, error) {
	t = strings.TrimSpace(t)

	if t == "" {
		return "", fmt.Errorf("empty title field")
	}

	if !utils.IsAlphanumeric(t) {
		return "", fmt.Errorf("expects only letters and numbers")
	}

	if !(len(t) <= maxTitleLength) {
		return "", fmt.Errorf("title is max %v characters", maxTitleLength)
	}

	return Title(t), nil
}
