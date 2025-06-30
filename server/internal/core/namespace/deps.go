package namespace

import (
	"context"

	namespace_store "github.com/server/internal/adapters/store/postgres/namespace"
	"github.com/server/internal/core/namespace/domain"
	"github.com/server/pkg/utils"
)

type INameSpaceStore interface {
	Select(ctx context.Context, id utils.UUID) (*namespace_store.NamespaceModel, error)
	Selects(ctx context.Context) ([]namespace_store.NamespaceModel, error)
	Update(ctx context.Context, input domain.SanitizedNamespace) (*namespace_store.NamespaceModel, error)
	Upsert(ctx context.Context, input domain.SanitizedNamespace) (*namespace_store.NamespaceModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteSoft(ctx context.Context, id utils.UUID) (*namespace_store.NamespaceModel, error)
}

