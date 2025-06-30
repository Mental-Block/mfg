package relation

import (
	"context"

	relation_store "github.com/server/internal/adapters/store/postgres/relation"
	"github.com/server/internal/core/relation/domain"
	"github.com/server/pkg/utils"
)

type IRelationStore interface {
	Select(ctx context.Context, id utils.UUID) (*relation_store.RelationModel, error)
	Selects(ctx context.Context) ([]relation_store.RelationModel, error)
	SelectsByFields(ctx context.Context, input domain.SanitizedRelation) ([]relation_store.RelationModel, error)
	Update(ctx context.Context, input domain.SanitizedRelation) (*relation_store.RelationModel, error)
	Upsert(ctx context.Context, input domain.SanitizedRelation) (*relation_store.RelationModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteSoft(ctx context.Context, id utils.UUID) (*relation_store.RelationModel, error)
}

type ISpiceRelationStore interface {
	Check(ctx context.Context, rel domain.Relation) (bool, error)
	BatchCheck(ctx context.Context, relations []domain.Relation) ([]domain.CheckPair, error)
	Delete(ctx context.Context, rel domain.Relation) error
	Add(ctx context.Context, rel domain.Relation) error
	LookupSubjects(ctx context.Context, rel domain.Relation) ([]string, error)
	LookupResources(ctx context.Context, rel domain.Relation) ([]string, error)
	ListRelations(ctx context.Context, rel domain.Relation) ([]domain.Relation, error)
}

