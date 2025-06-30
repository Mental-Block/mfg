package schema

import (
	_ "embed"

	"github.com/google/uuid"

	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	relation "github.com/server/internal/core/relation/domain"
	role "github.com/server/internal/core/role/domain"
)

//go:embed base_schema.zed
var BaseSchemaZed string

// SpiceDB readable format is stored in predefined_schema.txt
const (
	// Global IDs
	PlatformId = "platform"
)

var (
	PlatformOrgID = uuid.Nil
)

const (
	// platform permissions
	PlatformSuperPermission 	permission.PermissionName = "superuser"
	PlatformCheckPermission 	permission.PermissionName = "check"

	// permissions
	ListPermission              permission.PermissionName = "list"
	GetPermission               permission.PermissionName = "get"
	CreatePermission            permission.PermissionName = "create"
	UpdatePermission            permission.PermissionName = "update"
	DeletePermission            permission.PermissionName = "delete"
	RoleManagePermission        permission.PermissionName = "rolemanage"
	PolicyManagePermission      permission.PermissionName = "policymanage"
	ProjectListPermission       permission.PermissionName = "projectlist"
	GroupListPermission         permission.PermissionName = "grouplist"
	ProjectCreatePermission     permission.PermissionName = "projectcreate"
	GroupCreatePermission       permission.PermissionName = "groupcreate"
	ResourceListPermission      permission.PermissionName = "resourcelist"
	InvitationListPermission    permission.PermissionName = "invitationlist"
	InvitationCreatePermission  permission.PermissionName = "invitationcreate"
	AcceptPermission            permission.PermissionName = "accept"
	ServiceUserManagePermission permission.PermissionName = "serviceusermanage"
	ManagePermission            permission.PermissionName = "manage"
	
	// synthetic permission
	MembershipPermission permission.PermissionName = "membership"
)

const (
  // relations 
	PlatformRelationName     relation.RelationName = "platform"
	AdminRelationName        relation.RelationName = "admin"
	OrganizationRelationName relation.RelationName = "organization"
	UserRelationName         relation.RelationName = "user"
	ProjectRelationName      relation.RelationName = "project"
	GroupRelationName        relation.RelationName = "group"
	MemberRelationName       relation.RelationName = "member"
	OwnerRelationName        relation.RelationName = "owner"
	RoleRelationName         relation.RelationName = "role"
	RoleGrantRelationName    relation.RelationName = "granted"
	RoleBearerRelationName   relation.RelationName = "bearer"
)

const (
	// principals
	UserPrincipal        namespace.NameSpaceName = "app/user"
	ServiceUserPrincipal namespace.NameSpaceName = "app/serviceuser"
	GroupPrincipal       namespace.NameSpaceName = "app/group"
	SuperUserPrincipal   namespace.NameSpaceName = "app/superuser"
)

// DefaultNamespace is the default namespace for predefined entities
const ( 
	DefaultNamespace 	  namespace.NameSpaceName = "app"
	PlatformNamespace     namespace.NameSpaceName = "app/platform"
	OrganizationNamespace namespace.NameSpaceName = "app/organization"
	ProjectNamespace      namespace.NameSpaceName = "app/project"
	GroupNamespace        namespace.NameSpaceName = "app/group"
	RoleBindingNamespace  namespace.NameSpaceName = "app/rolebinding"
	RoleNamespace         namespace.NameSpaceName = "app/role"
	InvitationNamespace   namespace.NameSpaceName = "app/invitation"
)

const (
	// Roles
	RoleOrganizationAdminister role.RoleName = "app_organization_administer"
	RoleOrganizationViewer  role.RoleName = "app_organization_viewer"
	RoleOrganizationManager role.RoleName = "app_organization_manager"
	RoleOrganizationOwner   role.RoleName = "app_organization_owner"

	RoleProjectAdminister role.RoleName = "app_project_administer"
	RoleProjectOwner   role.RoleName = "app_project_owner"
	RoleProjectManager role.RoleName = "app_project_manager"
	RoleProjectViewer  role.RoleName = "app_project_viewer"

	GroupOwnerRole  role.RoleName = "app_group_owner"
	GroupMemberRole role.RoleName = "app_group_member"
)

// check if string provided is a system level namespace
func IsSystemSchema(ns string) bool {
	test := namespace.NameSpaceName(ns)

	switch (test) {
		case PlatformNamespace: 
			return true
		case OrganizationNamespace: 
			return true
		case ProjectNamespace: 
			return true
		case UserPrincipal: 
			return true
		case ServiceUserPrincipal: 
			return true
		case SuperUserPrincipal: 
			return true
		case GroupPrincipal: 
			return true
		default:		
			return false
	}
}