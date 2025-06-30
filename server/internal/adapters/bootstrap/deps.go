package bootstrap

import (
	"context"

	"github.com/server/internal/adapters/bootstrap/schema"
	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	role "github.com/server/internal/core/role/domain"
)

type INamespaceService interface {
	Upsert(ctx context.Context, ns namespace.Namespace) (*namespace.Namespace, error)
}

type IPermissionService interface {
	Upsert(ctx context.Context, perm permission.Permission) (*permission.Permission, error) 
	List(ctx context.Context) ([]permission.Permission, error)
}

type IUserService interface {
	NewSuper(ctx context.Context, id string, relation string) error
}

type IRoleService interface {
	Get(ctx context.Context, id string) (*role.Role, error)
	Upsert(ctx context.Context, toCreate role.Role) (*role.Role, error)
}

type IFileService interface {
	GetDefinition(ctx context.Context) (*schema.ServiceDefinition, error)
}

type IAuthzEngine interface {
	WriteSchema(ctx context.Context, schema string) error
}
