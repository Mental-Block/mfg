package schema

import (
	"strings"

	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	"github.com/server/pkg/utils"
)

// ResourcePermission with which roles will be created. Whenever an action is performed
// subject access permissions are checked with subject required permissions
type ResourcePermission struct {
	// simple name
	Name permission.PermissionName

	// Namespace is an object over which authz rules will be applied
	Namespace   namespace.NameSpaceName
	Description string

	// Key is a unique identifier composed of namespace and verb
	// for example: "app.platform.list" which is composed as service.resource.verb
	// here app.platform is namespace and list is name of the permission
	Key PermissionKey
}

func (r ResourcePermission) GetName() permission.PermissionName {
	if r.Name != "" {
		return r.Name
	}
	
	_, name := r.Key.GetNamespaceAndVerb()

	return name
}

func (r ResourcePermission) GetNamespace() namespace.NameSpaceName {
	if r.Namespace != "" {
		return r.Namespace
	}

	namespace, _ := r.Key.GetNamespaceAndVerb()

	return namespace
}

func (r ResourcePermission) Slug() string {
	if r.Key == "" {
		return r.Namespace.BuildRelation(r.Name.String())
	}
	
	namespace, name := r.Key.GetNamespaceAndVerb()

	return namespace.BuildRelation(name.String())
}

type PermissionKey string

// converts permission key back to its underlying type
func (s PermissionKey) String() string {
	return string(s)
}

// Gets the namespace, resource and verb the key is associated with   
// example app_platform_list -> app/platform, list
func (p PermissionKey) GetNamespaceAndVerb() (namespace.NameSpaceName, permission.PermissionName) {
	parts := strings.Split(p.String(), utils.RelationDilimeter)

	if len(parts) != 3 {
		return "", ""
	}

	namespace := namespace.BuildNamespace(parts[0], parts[1])

	return namespace,  permission.PermissionName(parts[2])
}
