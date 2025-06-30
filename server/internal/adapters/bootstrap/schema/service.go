package schema

import (
	role "github.com/server/internal/core/role/domain"
)

// ServiceDefinition is provided by user for a service
type ServiceDefinition struct {
	Roles       []RoleDefinition     `yaml:"roles"`
	Permissions []ResourcePermission `yaml:"permissions"`
}

// MergeServiceDefinitions merges multiple service definitions into one
// and deduplicate roles and permissions
func MergeServiceDefinitions(definitions ...ServiceDefinition) *ServiceDefinition {
	
	roles := make(map[role.RoleName]RoleDefinition)
	
	permissions := make(map[string]ResourcePermission)

	for _, definition := range definitions {

		for _, role := range definition.Roles {
			roles[role.Name] = role
		}

		for _, permission := range definition.Permissions {
			permissions[permission.Slug()] = permission
		}
	}

	roleList := make([]RoleDefinition, 0, len(roles))
	
	for _, role := range roles {
		roleList = append(roleList, role)
	}

	permissionList := make([]ResourcePermission, 0, len(permissions))
	
	for _, permission := range permissions {
		permissionList = append(permissionList, permission)
	}

	return &ServiceDefinition{
		Roles:       roleList,
		Permissions: permissionList,
	}
}
