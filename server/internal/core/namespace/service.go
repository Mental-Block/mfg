package namespace

import (
	"context"

	"github.com/server/internal/core/namespace/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of INameSpaceService should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IUserService interface {
 	Get(ctx context.Context, id string) (*domain.Namespace, error)
	Upsert(ctx context.Context, ns domain.Namespace) (*domain.Namespace, error)
	Update(ctx context.Context, ns domain.Namespace) (*domain.Namespace, error)
	List(ctx context.Context) ([]domain.Namespace, error)
}

type NameSpaceService struct {
	store INameSpaceStore
}

func NewService(store INameSpaceStore) *NameSpaceService {
	return &NameSpaceService{
		store: store,
	}
}

func (s NameSpaceService) Get(ctx context.Context, id string) (*domain.Namespace, error) {

	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	models, err := s.store.Select(ctx, UUID)

	if err != nil {
		return nil, err
	}

	sanitized, err := models.Transform()
	
	if (err != nil) {
		return nil, err
	}

	namespace := sanitized.Transform()

	return namespace, nil
}

func (s NameSpaceService) Upsert(ctx context.Context, ns domain.Namespace) (*domain.Namespace, error) {

	santized, err := ns.Transform()

	if (err != nil) {
		return nil, err
	}

	models, err := s.store.Upsert(ctx, *santized)

	if err != nil {
		return nil, err
	}

	sanitizedNamespace, err := models.Transform()
	
	if (err != nil) {
		return nil, err
	}

	namespace := sanitizedNamespace.Transform()

	return namespace, nil
}

func (s NameSpaceService) List(ctx context.Context) ([]domain.Namespace, error) {

	models, err := s.store.Selects(ctx)
	
	if err != nil {
		return nil, err
	}

	namespaces := make([]domain.Namespace, len(models))

	for i := range namespaces {

		santized, err :=  models[i].Transform()

		if (err != nil) {
			return nil, err
		}

		namespaces[i] = *santized.Transform()
	}

	return namespaces, nil
}

func (s NameSpaceService) Update(ctx context.Context, ns domain.Namespace) (*domain.Namespace, error) {
	santized, err := ns.Transform()

	if (err != nil) {
		return nil, err
	}

	models, err := s.store.Update(ctx, *santized)

	sanitizedNamespace, err := models.Transform()
	
	if (err != nil) {
		return nil, err
	}

	namespace := sanitizedNamespace.Transform()

	return namespace, nil
}
