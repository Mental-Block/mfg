package user

import (
	"context"

	user_store "github.com/server/internal/adapters/store/postgres/user"
	auth_domain "github.com/server/internal/core/auth/domain"
	relation_domain "github.com/server/internal/core/relation/domain"
	"github.com/server/internal/core/user/domain"
	"github.com/server/pkg/utils"
)

type IRelationService interface {
	BatchCheckPermission(ctx context.Context, relations []relation_domain.Relation) ([]relation_domain.CheckPair, error)
	LookupSubjects(ctx context.Context, rel relation_domain.Relation) ([]string, error)
	LookupResources(ctx context.Context, rel relation_domain.Relation) ([]string, error)
	New(ctx context.Context, input relation_domain.Relation) (*relation_domain.Relation, error)
 	Remove(ctx context.Context, rel relation_domain.Relation) error
}

type IUserStore interface {
	Delete(ctx context.Context, id utils.UUID) (*utils.UUID, error) 
	Select(ctx context.Context, id utils.UUID) (*user_store.UserModel, error)
	Insert(ctx context.Context, input domain.SanitizedUser) (*user_store.UserModel, error) 
	Selects(ctx context.Context) ([]user_store.UserModel, error)
	SelectByEmail(ctx context.Context, email auth_domain.Email) (*user_store.UserModel, error) 
	Update(ctx context.Context, input domain.SanitizedUser) (*user_store.UserModel, error)
}
