package schema

import (
	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	role "github.com/server/internal/core/role/domain"
)

// RoleDefinition are a set of permissions which can be assigned to a user or group
type RoleDefinition struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Name        role.RoleName`yaml:"name"`
	Scopes      []namespace.NameSpaceName `yaml:"scopes"`
	Permissions []permission.PermissionName `yaml:"permissions"`
}

var PredefinedRoles = []RoleDefinition{
	// org
	{
		Title: "Organization Owner",
		Name:  RoleOrganizationOwner,
		Permissions: []permission.PermissionName{
			"app_organization_administer",
		},
		Scopes: []namespace.NameSpaceName{
			OrganizationNamespace,
		},
	},
	{
		Title: "Organization Manager",
		Name:  RoleOrganizationManager,
		Permissions: []permission.PermissionName{
			"app_organization_update",
			"app_organization_get",
			"app_organization_projectcreate",
			"app_organization_projectlist",
			"app_organization_groupcreate",
			"app_organization_grouplist",
			"app_organization_serviceusermanage",
			"app_project_get",
			"app_project_update",
		},
		Scopes: []namespace.NameSpaceName{
			OrganizationNamespace,
		},
	},
	{
		Title: "Organization Access Manager",
		Name:  "app_organization_accessmanager",
		Permissions: []permission.PermissionName{
			"app_organization_invitationcreate",
			"app_organization_invitationlist",
			"app_organization_rolemanage",
			"app_organization_policymanage",
		},
		Scopes: []namespace.NameSpaceName{
			OrganizationNamespace,
		},
	},
	{
		Title: "Organization Viewer",
		Name:  RoleOrganizationViewer,
		Permissions: []permission.PermissionName{
			"app_organization_get",
		},
		Scopes: []namespace.NameSpaceName{
			OrganizationNamespace,
		},
	},
	{
		Title: "Organization Group Viewer",
		Name:  RoleOrganizationViewer,
		Permissions: []permission.PermissionName{
			"app_organization_get",
		},
		Scopes: []namespace.NameSpaceName{
			OrganizationNamespace,
		},
	},
	{
		Title: "Project Owner",
		Name:  RoleProjectOwner,
		Permissions: []permission.PermissionName{
			"app_project_administer",
		},
		Scopes: []namespace.NameSpaceName{
			ProjectNamespace,
		},
	},
	{
		Title: "Project Manager",
		Name:  RoleProjectManager,
		Permissions: []permission.PermissionName{
			"app_project_update",
			"app_project_get",
			"app_project_resourcelist",
			"app_organization_projectcreate",
			"app_organization_projectlist",
			"app_organization_grouplist",
		},
		Scopes: []namespace.NameSpaceName{
			ProjectNamespace,
		},
	},
	{
		Title: "Project Viewer",
		Name:  RoleProjectViewer,
		Permissions: []permission.PermissionName{
			"app_project_get",
		},
		Scopes: []namespace.NameSpaceName{
			ProjectNamespace,
		},
	},
	{
		Title: "Group Owner",
		Name:  GroupOwnerRole,
		Permissions: []permission.PermissionName{
			"app_group_administer",
		},
		Scopes: []namespace.NameSpaceName{
			GroupNamespace,
		},
	},
	{
		Title: "Group Member",
		Name:  GroupMemberRole,
		Permissions: []permission.PermissionName{
			"app_group_get",
		},
		Scopes: []namespace.NameSpaceName{
			GroupNamespace,
		},
	},
}
