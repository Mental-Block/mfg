package bootstrap

import (
	"context"
	"fmt"

	"github.com/server/internal/adapters/bootstrap/schema"
	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	role "github.com/server/internal/core/role/domain"

	"github.com/authzed/spicedb/pkg/caveats/types"
	"github.com/server/pkg/utils"

	azcore "github.com/authzed/spicedb/pkg/proto/core/v1"
)

var (
	defaultOrgID = schema.PlatformOrgID.String()
)

type Service struct {
	config       	  Config
	schemaConfig      IFileService
	namespaceService  INamespaceService
	roleService       IRoleService
	permissionService IPermissionService
	authzEngine       IAuthzEngine
	userService       IUserService
}

func NewBootstrapService(
	config Config,
	schemaConfig IFileService,
	namespaceService INamespaceService,
	roleService IRoleService,
	actionService IPermissionService,
	userService IUserService,
	authzEngine IAuthzEngine,
) *Service {
	return &Service{
		config:       	   config,
		schemaConfig:      schemaConfig,
		namespaceService:  namespaceService,
		roleService:       roleService,
		permissionService: actionService,
		userService:       userService,
		authzEngine:       authzEngine,
	}
}

func (s Service) MigrateSchema(ctx context.Context) error {
	customServiceDefinition, err := s.schemaConfig.GetDefinition(ctx)
	
	if err != nil {
		return err
	}

	return s.AppendSchema(ctx, *customServiceDefinition)
}

func (s Service) AppendSchema(ctx context.Context, customServiceDefinition schema.ServiceDefinition) error {
	// get existing permissions and append to the new definition
	// this is required to avoid overriding existing permissions in authzed engine
	var existingServiceDefinition schema.ServiceDefinition

	existingPermissions, err := s.permissionService.List(ctx)

	if err != nil {
		return nil
	}

	for _, existingPermission := range existingPermissions {
		
		description := ""
		if existingPermission.Metadata != nil {
			if v, ok := existingPermission.Metadata["description"]; !ok {
				description = v.(string)
			}
		}
		// // simple name
		// Name domain.Permission

		// // Namespace is an object over which authz rules will be applied
		// Namespace   domain.NameSpace
		// Description string

		// // Key is a unique identifier composed of namespace and name
		// // for example: "app.platform.list" which is composed as service.resource.verb
		// // here app.platform is namespace and list is name of the permission
		// Key domain.PermissionKey

		existingServiceDefinition.Permissions = append(existingServiceDefinition.Permissions, schema.ResourcePermission{
			Name: permission.PermissionName(existingPermission.Name),
			Namespace: namespace.NameSpaceName(existingPermission.NamespaceId),
			Description: description,
		})
	}

	return s.applySchema(ctx, schema.MergeServiceDefinitions(customServiceDefinition, existingServiceDefinition))
}

// applySchema builds and apply schema over az engine and db
// schema is composed of inbuilt definitions and custom user defined services
// this is idempotent operation and overrides existing schema
func (s Service) applySchema(ctx context.Context, customServiceDefinition *schema.ServiceDefinition) error {
	var err error

	// filter out default app namespace permissions
	var filteredPermissions []schema.ResourcePermission

	for _, permission := range customServiceDefinition.Permissions {

		if namespace, _ := permission.Namespace.Split(); namespace != schema.DefaultNamespace.String() {
			filteredPermissions = append(filteredPermissions, permission)
		}
	}

	customServiceDefinition.Permissions = filteredPermissions

	// build az schema with user defined services
	authzedDefinitions := GetBaseAZSchema()

	// TODO: GENERATE caveatTypeSet IN ApplyServiceDefinitionOverAZSchema. 
	// right now we cannot use caveats in DSL to convert it back to scema deff
	caveatTypeSet := &types.TypeSet{} 

	authzedDefinitions, err = ApplyServiceDefinitionOverAZSchema(customServiceDefinition, authzedDefinitions)

	if err != nil {
		return utils.NewErrorf(utils.ErrorCodeUnknown, "MigrateSchema: error applying schema over base: %s", err)
	}

	// validate prepared az schema
	authzedSchemaSource, err := PrepareSchemaAsAZSource(authzedDefinitions, caveatTypeSet)

	if err != nil {
		return fmt.Errorf("PrepareSchemaAsAZSource: %w", err)
	}

	if err = ValidatePreparedAZSchema(ctx, authzedSchemaSource); err != nil {
		return fmt.Errorf("ValidatePreparedAZSchema: %w", err)
	}

	// apply app to db
	appServiceDefinition, err := BuildServiceDefinitionFromAZSchema(authzedDefinitions)
	
	if err != nil {
		return fmt.Errorf("BuildServiceDefinitionFromAZSchema : %w", err)
	}

	if err = s.migrateAZDefinitionsToDB(ctx, authzedDefinitions); err != nil {
		return fmt.Errorf("migrateAZDefinitionsToDB : %w", err)
	}

	if err = s.migrateServiceDefinitionToDB(ctx, appServiceDefinition); err != nil {
		return fmt.Errorf("migrateServiceDefinitionToDB : %w", err)
	}

	// apply azSchema to engine
	if err = s.authzEngine.WriteSchema(ctx, authzedSchemaSource); err != nil {
		return fmt.Errorf("%w: %s", schema.ErrMigration, err.Error())
	}

	return nil
}

// MakeSuperUsers promote ordinary users to superuser
func (s Service) MakeSuperUsers(ctx context.Context) error {
	for _, userID := range s.config.SuperUsers {
		if err := s.userService.NewSuper(ctx, userID, schema.AdminRelationName.String()); err != nil {
			return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to make super users")
		}
	}
	
	return nil
}

// MigrateRoles migrate predefined roles to org
func (s Service) MigrateRoles(ctx context.Context) error {
	var err error

	// migrate predefined roles to org
	for _, defRole := range schema.PredefinedRoles {
		if err = s.createRole(ctx, defaultOrgID, defRole); err != nil {
			return err
		}
	}

	// migrate user defined roles to org
	serviceDefinition, err := s.schemaConfig.GetDefinition(ctx)
	
	if err != nil {
		return err
	}

	for _, defRole := range serviceDefinition.Roles {
		if err = s.createRole(ctx, defaultOrgID, defRole); err != nil {
			return err
		}
	}
	return nil
}

func (s Service) createRole(ctx context.Context, orgID string, defRole schema.RoleDefinition) error {
	
	if _, err := s.roleService.Get(ctx, defRole.Name.String()); err == nil {
		// role already exists
		return nil
	}
	
	scopes := make([]string, len(defRole.Scopes))

	for i, scope := range defRole.Scopes {
		scopes[i] = scope.String()
	}

	permissions := make([]string, len(defRole.Permissions)) 

	for i, perm := range defRole.Permissions {
		permissions[i] = perm.String()
	}

	_, err := s.roleService.Upsert(ctx, role.Role{
		Title:       defRole.Title,
		Name:        defRole.Name.String(),
		OrgId:       orgID,
		Permissions: permissions,
		Scopes:      scopes,
		Metadata: map[string]any{
			"description": defRole.Description,
		},
		Active: role.Enabled.Bool(),
	})

	if err != nil {
		return fmt.Errorf("can't migrate role: %w: %s", schema.ErrMigration, err.Error())
	}

	return nil
}

func (s Service) migrateServiceDefinitionToDB(ctx context.Context, appServiceDefinition schema.ServiceDefinition) error {
	
	// iterate over definition resources
	for _, perm := range appServiceDefinition.Permissions {
		
		// create permissions if needed
		_, err := s.permissionService.Upsert(ctx, permission.Permission{
			Name:        perm.GetName().String(),
			NamespaceId: perm.GetNamespace().String(),
			Metadata: map[string]any{
				"description": perm.Description,
			},
		})

		if err != nil {
			return fmt.Errorf("permissionService.Upsert: %s: %w", err.Error(), schema.ErrMigration)
		}

	}

	return nil
}

// migrateAZDefinitionsToDB will ensure wll the namespaces are already created in database which will be used
// throughout the application
func (s Service) migrateAZDefinitionsToDB(ctx context.Context, azDefinitions []*azcore.NamespaceDefinition) error {
	// iterate over all az definitions and convert frontier namespace
	for _, azDef := range azDefinitions {
		
		// create namespace if needed
		_, err := s.namespaceService.Upsert(ctx, namespace.Namespace{
			Name: azDef.GetName(),
		})

		if err != nil {
			return fmt.Errorf("namespaceService.Upsert: %w: %s", schema.ErrMigration, err.Error())
		}

	}

	return nil
}
