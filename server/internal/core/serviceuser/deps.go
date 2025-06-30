package serviceuser

import (
	"context"

	serviceuser_store "github.com/server/internal/adapters/store/postgres/serviceuser"
	relation "github.com/server/internal/core/relation/domain"
	"github.com/server/internal/core/serviceuser/domain"
	"github.com/server/pkg/utils"
)

type IRelationService interface {
	New(ctx context.Context, input relation.Relation) (*relation.Relation, error)
 	Remove(ctx context.Context, rel relation.Relation) error
	LookupSubjects(ctx context.Context, rel relation.Relation) ([]string, error)
	CheckPermission(ctx context.Context, rel relation.Relation) (bool, error)
	BatchCheckPermission(ctx context.Context, relations []relation.Relation) ([]relation.CheckPair, error)
}

type IServiceUserStore interface {
	Insert(ctx context.Context, input domain.SanitizedServiceUser) (*serviceuser_store.ServiceUserModel, error)
	Selects(ctx context.Context, filter domain.ServiceUserFilter) ([]serviceuser_store.ServiceUserModel, error) 
	SelectByIds(ctx context.Context, ids []utils.UUID) ([]serviceuser_store.ServiceUserModel, error)
	Select(ctx context.Context, id utils.UUID) (*serviceuser_store.ServiceUserModel, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	DeleteUserAndCredentials(ctx context.Context, id utils.UUID) error
}

type IServiceUserCredentialStore interface {
	Select(ctx context.Context, id utils.UUID)(*serviceuser_store.ServiceUserCredentialModel, error)
	Selects(ctx context.Context, filter domain.CredentialFilter)([]serviceuser_store.ServiceUserCredentialModel, error)
 	Insert(ctx context.Context, input domain.SanitizedServiceUserCredential)(*serviceuser_store.ServiceUserCredentialModel, error)
	DeleteByUserId(ctx context.Context, id utils.UUID) (*utils.UUID, error)
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error)
}