package bootstrap

import (
	"context"
	"fmt"
	"strings"

	"github.com/authzed/spicedb/pkg/caveats/types"
	aznamespace "github.com/authzed/spicedb/pkg/namespace"
	azcore "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
	"github.com/authzed/spicedb/pkg/schemadsl/generator"
	"github.com/authzed/spicedb/pkg/schemautil"

	"github.com/server/internal/adapters/bootstrap/schema"
	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	"github.com/server/pkg/utils"
)

func ValidatePreparedAZSchema(ctx context.Context, azSchemaSource string) error {
	// compile and validate generated schema
	tenantName := "auth"

	updatedSchema, err := compiler.Compile(compiler.InputSchema{
		Source:       "generated",
		SchemaString: azSchemaSource,
	}, compiler.ObjectTypePrefix(tenantName))

	if err != nil {
		return fmt.Errorf("compile: failed to compile authz schema: %w", err)
	}

	if _, err = schemautil.ValidateSchemaChanges(ctx, updatedSchema, false); err != nil {
		return fmt.Errorf("ValidateSchemaChanges: failed to validate authz schema: %w", err)
	}

	return nil
}

func PrepareSchemaAsAZSource(authzedDefinitions []*azcore.NamespaceDefinition, caveatTypeSet *types.TypeSet) (string, error) {
	preparedSchemaString := ""
	
	for _, def := range authzedDefinitions {
		generatedDefString, _, err := generator.GenerateSource(def, caveatTypeSet)
		
		if err != nil {
			return "", fmt.Errorf("generateSource: failed to compile authz schema: %w", err)
		}

		preparedSchemaString = fmt.Sprintf("%s\n\n%s", preparedSchemaString, generatedDefString)
	}

	return preparedSchemaString, nil
}

func GetBaseAZSchema() []*azcore.NamespaceDefinition {
	tenantName := "auth"
	
	compiledSchema, err := compiler.Compile(compiler.InputSchema{
		Source:       "base_schema.zed",
		SchemaString: schema.BaseSchemaZed,
	}, compiler.ObjectTypePrefix(tenantName))
	
	if err != nil {
		// this should not happen
		panic(err)
	}
	
	return compiledSchema.ObjectDefinitions
}

// BuildServiceDefinitionFromAZSchema converts authzed schema to auth service definition.
// This conversion is lossy, and it only keeps list of permissions used in the schema per resource
func BuildServiceDefinitionFromAZSchema(azDefinitions []*azcore.NamespaceDefinition, serviceFilter ...string) (schema.ServiceDefinition, error) {
	
	resourcePermissions := []schema.ResourcePermission{}
	
	// iterate over namespace to find services and permissions
	for _, def := range azDefinitions {
		if def.GetName() == string(schema.RoleBindingNamespace) {
			// build permission set for all namespaces using roles to bind themselves
			for _, rel := range def.GetRelation() {
				if rel.GetUsersetRewrite() != nil { // not nil for permissions in zed file
					
					permissionParts := strings.Split(rel.GetName(), utils.RelationDilimeter)
					
					var service, resource, perm string
					
					switch len(permissionParts) {
					case 3:
						service, resource, perm = permissionParts[0], permissionParts[1], permissionParts[2]
					case 2:
						perm, resource = permissionParts[1], permissionParts[0]
					case 1:
						perm = permissionParts[0]
					default:
						service, resource = permissionParts[0], permissionParts[1]
						perm = strings.Join(permissionParts[2:], "")
					}

					if len(serviceFilter) > 0 && !utils.Contains(serviceFilter, service) {
						// ignore service if filter was requested, and it doesn't match
						continue
					}

					resourcePermissions = append(resourcePermissions, schema.ResourcePermission{
						Name:      permission.PermissionName(perm),
						Namespace: namespace.BuildNamespace(service, resource),
					})
				}
			}
		}
	}

	return schema.ServiceDefinition{
		Permissions: resourcePermissions,
	}, nil
}

// ApplyServiceDefinitionOverAZSchema applies the provided user defined service over existing schema
// and returns the updated schema
func ApplyServiceDefinitionOverAZSchema(serviceDef *schema.ServiceDefinition, existingDefinitions []*azcore.NamespaceDefinition) ([]*azcore.NamespaceDefinition, error) {
	// keep relations/permissions required to be appended in existing definitions
	// this is required to bind roles over application authz hierarchy
	var relationsForOrg []*azcore.Relation
	var relationsForProject []*azcore.Relation
	var relationsForRole []*azcore.Relation
	var relationsForRoleBinding []*azcore.Relation

	// gather list of resources
	namespaceNameToPermissionNameMap := map[namespace.NameSpaceName][]string{}
	for _, perm := range serviceDef.Permissions {
		namespaceNameToPermissionNameMap[perm.GetNamespace()] = append(
			namespaceNameToPermissionNameMap[perm.GetNamespace()], 
			perm.GetName().String(),
		)
	}

	// prepare new definition with its own relations and permissions
	// and relations that need to be added in base definitions like org/project
	userDefinedAZServiceDefinitions := make([]*azcore.NamespaceDefinition, 0, len(namespaceNameToPermissionNameMap))
	for namespaceName, permissions := range namespaceNameToPermissionNameMap {
		var relationsForResource []*azcore.Relation
		for _, name := range permissions {
			
			relation := namespaceName.BuildRelation(name)

			// create permissions
			{
				// for resource
				nsRel, err := aznamespace.Relation(name, aznamespace.Union(
					aznamespace.ComputedUserset(schema.OwnerRelationName.String()),
					aznamespace.TupleToUserset(schema.ProjectRelationName.String(), schema.RoleProjectAdminister.String()),
					aznamespace.TupleToUserset(schema.ProjectRelationName.String(), relation),
					aznamespace.TupleToUserset(schema.RoleGrantRelationName.String(), relation),
				), nil)
				

				if err != nil {
					return nil, err
				}

				relationsForResource = append(relationsForResource, nsRel)
			}
			{
				// for org
				nsRel, err := aznamespace.Relation(relation, aznamespace.Union(
					aznamespace.ComputedUserset(schema.OwnerRelationName.String()),
					aznamespace.TupleToUserset(schema.PlatformRelationName.String(), schema.PlatformSuperPermission.String()),
					aznamespace.TupleToUserset(schema.RoleGrantRelationName.String(), schema.RoleOrganizationAdminister.String()),
					aznamespace.TupleToUserset(schema.RoleGrantRelationName.String(), relation),
				), nil)

				if err != nil {
					return nil, err
				}

				relationsForOrg = append(relationsForOrg, nsRel)
			}
			{
				// for project
				nsRel, err := aznamespace.Relation(relation, aznamespace.Union(
					aznamespace.TupleToUserset(schema.OrganizationRelationName.String(), relation),
					aznamespace.TupleToUserset(schema.RoleGrantRelationName.String(), schema.RoleProjectAdminister.String()),
					aznamespace.TupleToUserset(schema.RoleGrantRelationName.String(), relation),
				), nil)
				
				if err != nil {
					return nil, err
				}

				relationsForProject = append(relationsForProject, nsRel)
			}
			{
				// for rolebinding
				nsRel, err := aznamespace.Relation(relation, aznamespace.Intersection(
					aznamespace.ComputedUserset(schema.RoleBearerRelationName.String()),
					aznamespace.TupleToUserset(schema.RoleRelationName.String(), relation),
				), nil)

				if err != nil {
					return nil, err
				}

				relationsForRoleBinding = append(relationsForRoleBinding, nsRel)
			}
			{
				// for role
				nsRel, err := aznamespace.Relation(relation, nil,
					aznamespace.AllowedPublicNamespace(schema.UserPrincipal.String()),
					aznamespace.AllowedPublicNamespace(schema.ServiceUserPrincipal.String()),
				)

				if err != nil {
					return nil, err
				}

				relationsForRole = append(relationsForRole, nsRel)
			}
		}

		// prepare an owner relation
		// either we can attach each user who creates the resource with owner relation or
		// create an owner role and assign it to the user when the resource is created
		relationsForResource = append(relationsForResource, aznamespace.MustRelation(schema.OwnerRelationName.String(), nil,
			aznamespace.AllowedRelation(schema.UserPrincipal.String(), generator.Ellipsis),
			aznamespace.AllowedRelation(schema.ServiceUserPrincipal.String(), generator.Ellipsis)))
		
		// attach service to project
		relationsForResource = append(relationsForResource, aznamespace.MustRelation(schema.ProjectRelationName.String(), nil,
			aznamespace.AllowedRelation(schema.ProjectNamespace.String(), generator.Ellipsis)))
		
			// attach role binding to service
		relationsForResource = append(relationsForResource, aznamespace.MustRelation(schema.RoleGrantRelationName.String(), nil,
			aznamespace.AllowedRelation(schema.RoleBindingNamespace.String(), generator.Ellipsis)))

		// prepare a new az definition
		resourceDef := aznamespace.Namespace(namespaceName.String(), relationsForResource...)
		userDefinedAZServiceDefinitions = append(userDefinedAZServiceDefinitions, resourceDef)
	}

	// append new definition to existing list of definitions
	newSetOfDefinitions := append(existingDefinitions, userDefinedAZServiceDefinitions...)

	if len(relationsForOrg) > 0 {
		for _, baseDef := range newSetOfDefinitions {
			switch namespace.NameSpaceName(baseDef.GetName()) {
			case schema.OrganizationNamespace:
				// populate app/organization with service permissions to allow bounding service roles at org level
				baseDef.Relation = append(baseDef.GetRelation(), relationsForOrg...)
			case schema.ProjectNamespace:
				// populate app/project with service permissions to allow bounding service roles at project level
				baseDef.Relation = append(baseDef.GetRelation(), relationsForProject...)
			case schema.RoleBindingNamespace:
				// populate app/rolebinding with service relations to allow checking service roles with permissions
				baseDef.Relation = append(baseDef.GetRelation(), relationsForRoleBinding...)
			case schema.RoleNamespace:
				// populate app/role with service permissions to allow building service roles with permissions
				baseDef.Relation = append(baseDef.GetRelation(), relationsForRole...)
			}
		}
	}

	return newSetOfDefinitions, nil
}
