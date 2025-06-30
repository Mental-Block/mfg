package blob

import (
	"context"
	"sort"
	"testing"

	"github.com/server/internal/adapters/bootstrap/schema"
	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetSchema(t *testing.T) {

	cfg := Config{
		StoragePath: "file://testdata",
		StorageSecret: "",
	}

	testBucket, err := NewStore(context.Background(), cfg)

	assert.NoError(t, err)

	s := SchemaStore{
		store: testBucket,
	}

	def, err := s.GetDefinition(context.Background())
	assert.NoError(t, err)

	expectedMap := &schema.ServiceDefinition{
		Roles: []schema.RoleDefinition{
			{
				Name: "compute_order_manager",
				Permissions: []permission.PermissionName{
					"compute/order:delete",
					"compute/order:update",
					"compute/order:get",
					"compute/order:list",
					"compute/order:create",
				},
			},
			{
				Name: "compute_order_viewer",
				Permissions: []permission.PermissionName{
					"compute/order:list",
					"compute/order:get",
				},
			},
			{
				Name: "compute_order_owner",
				Permissions: []permission.PermissionName{
					"compute/order:delete",
					"compute/order:update",
					"compute/order:get",
					"compute/order:create",
				},
			},
		},
		Permissions: []schema.ResourcePermission{
			{
				Name:      "delete",
				Namespace: "compute/order",
			},
			{
				Name:      "update",
				Namespace: "compute/order",
			},
			{
				Name:      "get",
				Namespace: "compute/order",
			},
			{
				Name:      "list",
				Namespace: "compute/order",
			},
			{
				Name:      "create",
				Namespace: "compute/order",
			},
			{
				Name:      "delete",
				Namespace: "database/order",
			},
			{
				Name:      "update",
				Namespace: "database/order",
			},
			{
				Name:      "get",
				Namespace: "database/order",
			},
		},
	}

	sort.Slice(def.Roles, func(i, j int) bool {
		return def.Roles[i].Name < def.Roles[j].Name
	})

	sort.Slice(expectedMap.Roles, func(i, j int) bool {
		return expectedMap.Roles[i].Name < expectedMap.Roles[j].Name
	})

	sort.Slice(def.Permissions, func(i, j int) bool {
		name := namespace.NameSpaceName(def.Permissions[i].Namespace).BuildRelation(def.Permissions[i].Name.String())
		name2 := namespace.NameSpaceName(def.Permissions[j].Namespace).BuildRelation(def.Permissions[j].Name.String())
		return name < name2
	})

	sort.Slice(expectedMap.Permissions, func(i, j int) bool {
		name := namespace.NameSpaceName(expectedMap.Permissions[i].Namespace).BuildRelation(expectedMap.Permissions[i].Name.String())
		name2 := namespace.NameSpaceName(expectedMap.Permissions[j].Namespace).BuildRelation(expectedMap.Permissions[j].Name.String())
		return name < name2
	})

	assert.Equal(t, expectedMap, def)
}
