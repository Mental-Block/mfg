package domain

import (
	"strings"

	namespace "github.com/server/internal/core/namespace/domain"

	"github.com/server/pkg/utils"
)

type PermissionName string

// converts permission back to its underlying type
func (s PermissionName) String() string {
	return string(s)
}

// IsValidPermission checks if the provided name is a valid permission
func (s PermissionName) IsValid() error {
	name := strings.TrimSpace(s.String())

	if name == "" {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, "empty permission field")
	}

	if !utils.IsAlphanumeric(name) {
		return  utils.NewErrorf(utils.ErrorCodeInvalidArgument, "expects only letters and numbers")
	}

	return nil
}

// generates a slug from a namespace and permission name. equvilant to namespace.BuildRelation
// example: app/resource, view -> app_resource_permission
func (s PermissionName) BuildSlug(namepace namespace.NameSpaceName) PermissionSlug {
	return PermissionSlug(namepace.BuildRelation(s.String()))
}
