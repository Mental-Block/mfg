package permission

import (
	"context"

	permission_store "github.com/server/internal/adapters/store/postgres/permission"
	"github.com/server/internal/core/permission/domain"
	"github.com/server/pkg/utils"
)

type IPermissionStore interface {
	Select(ctx context.Context, id utils.UUID) (*permission_store.PermissionModel, error)
	SelectBySlug(ctx context.Context, slug string) (*permission_store.PermissionModel, error)
	Selects(ctx context.Context) ([]permission_store.PermissionModel, error)
	Upsert(ctx context.Context, input domain.SanitizedPermission) (*permission_store.PermissionModel, error)
	Update(ctx context.Context, input domain.SanitizedPermission) (*permission_store.PermissionModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) 
}