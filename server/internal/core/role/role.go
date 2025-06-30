package role

import (
	"context"
	"fmt"

	"github.com/server/internal/adapters/bootstrap/schema"
	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	relation "github.com/server/internal/core/relation/domain"
	"github.com/server/internal/core/role/domain"
	"github.com/server/pkg/utils"
)

/*
 High level overview of RoleService should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IRoleService interface {
	Upsert(ctx context.Context, role domain.Role) (*domain.Role, error)
	Update(ctx context.Context, role domain.Role) (*domain.Role, error)
	Get(ctx context.Context, id string) (*domain.Role, error)
	List(ctx context.Context, f domain.Filter) ([]domain.Role, error)
	Delete(ctx context.Context, id string) error
}

type RoleService struct {
	roleStore 		  IRoleStore
	relationService   IRelationService
	permissionService IPermissionService
}

func NewRoleService(
	roleStore IRoleStore, 
	relationService IRelationService,
	permissionService IPermissionService,
) *RoleService {
	return &RoleService {
		roleStore,
		relationService,
		permissionService,
	}
}

var principals = []namespace.NameSpaceName{
	schema.UserPrincipal,
	schema.ServiceUserPrincipal,
}

func (s RoleService) Upsert(ctx context.Context, role domain.Role) (*domain.Role, error) {

	sanitizedRole, err := role.Transfrom()

	if err != nil {
		return nil, err
	}

	for idx, permName := range sanitizedRole.Permissions {	
		perm, err := s.permissionService.Get(ctx, permName.String())

		if err != nil {
			return nil, fmt.Errorf("%s: %w", permName, err)
		}

		slug :=  permission.PermissionName(permName.BuildSlug(namespace.NameSpaceName(perm.NamespaceId)))

		sanitizedRole.Permissions[idx] = slug
	}

	roleModel, err := s.roleStore.Upsert(ctx, *sanitizedRole)

	if err != nil {
		return nil, err
	}

	sanitzedRole, err := roleModel.Transform()

	if err != nil {
		return nil, err
	}

	newRole := sanitzedRole.Transfrom()

	// create relation between role and permissions
	if err := s.createRolePermissionRelation(ctx, sanitizedRole.Id, sanitizedRole.Permissions); err != nil {
		return nil, err
	}

	return newRole, nil
}

func (s RoleService) Get(ctx context.Context, id string) (*domain.Role, error) {

	uuid, err := utils.ConvertStringToUUID(id)

	if err != nil {
		return nil, err
	}

	roleModel, err := s.roleStore.Select(ctx, uuid)

	if err != nil {
		return nil, err
	}

	sanitizedRole, err := roleModel.Transform()

	if err != nil {
		return nil, err
	}

	role := sanitizedRole.Transfrom()

	return role, err
}

// return roles  that are/werer created at when we bootstrap in the build dependancies phase. 
//func (s RoleService) GetSystem(ctx context.Context, id string) (*domain.Role, error) {

	//pass empty orgId will get roles created that we created at runtime. 
	
	// passing empty orgID will return roles created by system
	//	return s.roleStore.GetByName(ctx, "", id)
//}



func (s RoleService) List(ctx context.Context, f domain.Filter) ([]domain.Role, error) {
	
	roleModels, err := s.roleStore.Selects(ctx)
	
	if err != nil {
		return nil, err
	}

	roles := make([]domain.Role, len(roleModels))
	
	for i := range roles {

		sanitizedRole, err := roleModels[i].Transform()	

		if err != nil {
			return nil, err
		}

		role := sanitizedRole.Transfrom()

		roles[i] = *role
	}	
	
	return roles, nil
}

func (s RoleService) Update(ctx context.Context, role domain.Role) (*domain.Role, error) {
	
	sanitizedRole, err := role.Transfrom()

	if err != nil {
		return nil, err
	}

	for idx, permName := range sanitizedRole.Permissions {	
		perm, err := s.permissionService.Get(ctx, permName.String())

		if err != nil {
			return nil, fmt.Errorf("%s: %w", permName, err)
		}

		slug :=  permission.PermissionName(permName.BuildSlug(namespace.NameSpaceName(perm.NamespaceId)))

		sanitizedRole.Permissions[idx] = slug
	}

	// delete all OLD/existing relation between role and permissions
	if err := s.deleteRolePermissionRelation(ctx, sanitizedRole.Id); err != nil {
		return nil, err
	}

	// NEW/create relation between role and permissions
	if err := s.createRolePermissionRelation(ctx, sanitizedRole.Id, sanitizedRole.Permissions); err != nil {
		return nil, err
	}

	roleModel, err := s.roleStore.Update(ctx, *sanitizedRole)

	if err != nil {
		return nil, err
	}

	backToSantizedRole, err := roleModel.Transform()

	if err != nil {
		return nil, err
	}

	return backToSantizedRole.Transfrom(), nil
}

func (s RoleService) Delete(ctx context.Context, id string) error {

	uuid, err := utils.ConvertStringToUUID(id)

	if err != nil {
		return err
	}

	err = s.relationService.Remove(ctx, relation.Relation{
		Object: relation.Object{
			Id:        id,
			Namespace: schema.RoleNamespace.String(),
		},
	})
	
	if err != nil {
		return err
	}

	_, err = s.roleStore.Delete(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (s RoleService) createRolePermissionRelation(ctx context.Context, id utils.UUID, permissions []permission.PermissionName) error {
	// create relation between role and permissions
	// for example for each permission:
	// app/role:org_owner#organization_delete@app/user:*
	// app/role:org_owner#organization_update@app/user:*
	// this needs to be created for each type of principles
	
	// add both principals the the relation
	for _, perm := range permissions {
		for _, principal := range principals {
			
			_, err := s.relationService.New(ctx, relation.Relation{
				Object: relation.Object{
					Id:        id.String(),
					Namespace: schema.RoleNamespace.String(),
				},
				Subject: relation.Subject{
					Id:        relation.SubjectAll, // all principles who have role will have access to these permissions
					Namespace: principal.String(),
				},
				RelationName: perm.String(),
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s RoleService) deleteRolePermissionRelation(ctx context.Context, id utils.UUID) error {
	// delete relation between role and permissions
	// for example for each permission:
	// app/role:org_owner#organization_delete@app/user:*
	// app/role:org_owner#organization_update@app/user:*
	// this needs to be created for each type of principles

	for _, principal := range principals {
		err := s.relationService.Remove(ctx, relation.Relation{
			Object: relation.Object{
				Id:        id.String(),
				Namespace: schema.RoleNamespace.String(),
			},
			Subject: relation.Subject{
				Id:        relation.SubjectAll, // all principles who have role will have access
				Namespace: principal.String(),
			},
		})

		if err != nil {
			return err
		}
	}
	
	return nil
}