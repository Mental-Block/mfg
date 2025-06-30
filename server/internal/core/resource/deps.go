package resource

import (
	"context"

	resource_store "github.com/server/internal/adapters/store/postgres/resource"
	"github.com/server/internal/core/resource/domain"

	"github.com/server/pkg/utils"
)

type IResourceStore interface {
	Upsert(ctx context.Context, input domain.SanitizedResource) (*resource_store.ResourceModel, error)
	Update(ctx context.Context, input domain.SanitizedResource) (*resource_store.ResourceModel, error) 
	Select(ctx context.Context, id utils.UUID) (*resource_store.ResourceModel, error)
	SelectByURN(ctx context.Context, urn string) (*resource_store.ResourceModel, error)
	Selects(ctx context.Context) ([]resource_store.ResourceModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
}


type YAML struct {
	Name         string              `json:"name" yaml:"name"`
	Backend      string              `json:"backend" yaml:"backend"`
	ResourceType string              `json:"resource_type" yaml:"resource_type"`
	Actions      map[string][]string `json:"actions" yaml:"actions"`
}

type IResourceBlobStore interface {
	GetAll(ctx context.Context) ([]YAML, error)
}

