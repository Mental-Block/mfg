package user

import (
	"context"
	"errors"

	"github.com/server/internal/adapters/bootstrap/schema"
	auth "github.com/server/internal/core/auth/domain"
	permission "github.com/server/internal/core/permission/domain"
	relation "github.com/server/internal/core/relation/domain"
	"github.com/server/internal/core/user/domain"

	"github.com/server/pkg/utils"
)

/*
 High level overview of UserService should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IUserService interface {
 	GetByEmail(ctx context.Context, email string) (*domain.User, error) 
	List(ctx context.Context) ([]domain.User, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	Remove(ctx context.Context, id string) (*string, error) 
	New(ctx context.Context, input domain.User) (*domain.User, error)
	NewSuper(ctx context.Context, id string, relation string) error
	RemoveSuper(ctx context.Context, id string) error
	IsSupers(ctx context.Context, ids []string, relationPermission string) ([]relation.Relation, error)
	IsSuper(ctx context.Context, id string, relationPermission string) (bool, error)
}

type UserService struct {
	userStore IUserStore
	relationService IRelationService
}

func NewUserService(
	user IUserStore,
	relation IRelationService,
) UserService {
	return UserService{
		userStore: user,
		relationService: relation,
	}
}

func (s *UserService) New(ctx context.Context, input domain.User) (*domain.User, error) {

	input.Id =  utils.NewUUID().String()

	sanitizedUser, err := input.Transform()
	
	if err != nil {
		return nil, err
	}

	userModel, err := s.userStore.Insert(ctx, sanitizedUser)
    
	if err != nil {
		return nil, utils.WrapErrorf(err,utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	sanitizedUsr, err := userModel.Transform()

	if err != nil {
		return nil, utils.WrapErrorf(err,utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	user := sanitizedUsr.Transform()

	return user, nil
}

func (s *UserService) Remove(ctx context.Context, id string) (string, error) {
	
	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return "", utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	dataId, err := s.userStore.Delete(ctx, UUID)

	if err != nil {
		return "", utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	return string(*dataId), nil
}

func (s *UserService) Get(ctx context.Context, id string) (*domain.User, error) {

	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	model, err := s.userStore.Select(ctx, UUID)

	if err != nil {	
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	sanitizedUser, err := model.Transform()
	
	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	user := sanitizedUser.Transform()

	return user, nil
}

func (s *UserService) Update(ctx context.Context, input domain.User) (*domain.User, error) {

	validUser, err := input.Transform()

	if err != nil {
		return nil, err
	}
	
	model, err := s.userStore.Update(ctx, validUser)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	sanitizedUser, err := model.Transform()

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	user := sanitizedUser.Transform()

	return user, nil
}

func (s *UserService) List(ctx context.Context) ([]domain.User, error) {
	
	models, err := s.userStore.Selects(ctx)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	users := make([]domain.User, len(models))

	for i := range users {
		santizedUser, err := models[i].Transform()
		
		if (err != nil) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
		}

		users[i] = *santizedUser.Transform()
	}

	return users, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {

	validEmail, err := auth.Email(email).NewEmail()

	if err != nil {
		return nil,  utils.NewErrorf(utils.ErrorCodeInvalidArgument, "invalid email format: %s", err)
	}

	model, err := s.userStore.SelectByEmail(ctx, validEmail)

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	sanitizedUser, err := model.Transform()

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	user := sanitizedUser.Transform()

	return user, nil
}


// platform level user permissions 
// prefer using service user over direct super user privledge
func (s UserService) NewSuper(ctx context.Context, id string, rel string) error {

	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	currentUser, err := s.userStore.Select(ctx, UUID)
	
	if err != nil {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	var name relation.RelationName
	var permission permission.PermissionName

	switch rel {
		case schema.MemberRelationName.String():
			name = schema.MemberRelationName
			permission = schema.PlatformCheckPermission
		case schema.AdminRelationName.String():
			name = schema.AdminRelationName
			permission = schema.PlatformSuperPermission
		default:
			return utils.WrapErrorf(
			errors.New("not a platform relation"), 
			utils.ErrorCodeInvalidArgument, 
			"possible options include: %s, %s",
			schema.AdminRelationName,
			schema.MemberRelationName,
		)
	}

	if ok, err := s.IsSuper(ctx, currentUser.Id.String(), permission.String()); err != nil {
		return err
	} else if ok {
		return nil
	}

	_, err = s.relationService.New(ctx, relation.Relation{
		Object: relation.Object{
			Id:        schema.PlatformId,
			Namespace: schema.PlatformNamespace.String(),
		},
		Subject: relation.Subject{
			Id:        currentUser.Id.String(),
			Namespace: schema.UserPrincipal.String(),
		},
		RelationName: name.String(),
	})
	
	return err
}

// removes user platform permissions
func (s UserService) RemoveSupper(ctx context.Context, id string) error {
	
	UUID, err := utils.ConvertStringToUUID(id) 
	
	if err != nil {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	currentUser, err := s.userStore.Select(ctx, UUID)
	
	if err != nil {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	// to check if the user has member relation, we need to check if the user has `check` permission on platform
	if ok, err := s.IsSuper(ctx, currentUser.Id.String(), schema.PlatformCheckPermission.String()); err != nil {
		return err
	} else if !ok {
		return nil
	}

	err = s.relationService.Remove(ctx, relation.Relation{
		Object: relation.Object{
			Id:        schema.PlatformId,
			Namespace: schema.PlatformNamespace.String(),
		},
		Subject: relation.Subject{
			Id:        currentUser.Id.String(),
			Namespace: schema.UserPrincipal.String(),
		},
		RelationName: schema.MemberRelationName.String(),
	})
	return err
}

// checks user platform permissions.
func (s UserService) IsSuper(ctx context.Context, id string, relationName string) (bool, error) {
	
	status, err := s.IsSupers(ctx, []string{ id }, relationName)

	if err != nil {
		return false, err
	}

	if (len(status) > 0) {
		return true, nil
	} 

	return false, nil
}

// check to see if a user can have relation to permissions  
func (s UserService) IsSupers(ctx context.Context, ids []string, relationName string) ([]relation.Relation, error) {

	relations := make([]relation.Relation, len(ids))

	for i, v := range ids {
		UUID, err := utils.ConvertStringToUUID(v) 
	
		if err != nil {
			return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
		}

		relations[i] = relation.Relation{
			Subject: relation.Subject{
				Id:        UUID.String(),
				Namespace: schema.UserPrincipal.String(),
			},
			Object: relation.Object{
				Id:        schema.PlatformId,
				Namespace: schema.PlatformNamespace.String(),
			},
			RelationName: relationName,
		}
	}

	statusForIds, err := s.relationService.BatchCheckPermission(ctx, relations)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrUserService)
	}

	successChecks := utils.Filter(statusForIds, func(pair relation.CheckPair) bool {
		return pair.Status
	})

	userRelations := utils.Map(successChecks, func(pair relation.CheckPair) relation.Relation {
		return pair.Relation
	})

	return userRelations, nil
}
