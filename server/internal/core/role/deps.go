package role

import (
	"context"

	"github.com/server/internal/adapters/store/postgres/role"
	permission "github.com/server/internal/core/permission/domain"
	relation "github.com/server/internal/core/relation/domain"
	"github.com/server/internal/core/role/domain"
	"github.com/server/pkg/utils"
)

type IRoleStore interface {
	Upsert(ctx context.Context, input domain.SanitizedRole) (*role.RoleModel, error)
	Update(ctx context.Context, input domain.SanitizedRole) (*role.RoleModel, error) 
	Select(ctx context.Context, id utils.UUID) (*role.RoleModel, error)
	Selects(ctx context.Context) ([]role.RoleModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteSoft(ctx context.Context, id utils.UUID) (*role.RoleModel, error) 
}

type IRelationService interface {	
	New(ctx context.Context, input relation.Relation) (*relation.Relation, error)
 	Remove(ctx context.Context, rel relation.Relation) error
}

type IPermissionService interface {
	Get(ctx context.Context, id string) (*permission.Permission, error)
}
