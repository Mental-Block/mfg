package domain

import (
	"strings"

	"github.com/server/pkg/utils"
)

type NameSpaceName string

var genericResourceType = "default"

// converts Namespace back to its underlying type
func (s NameSpaceName) String() string {
	return string(s)
}

// check if namespace is in a valid format
func (s NameSpaceName) IsValid() error {

	// empty values are valid namespaces
	if (s.String() == "") {
		return nil
	}
	
	service, resource := s.Split()

	// service and resource can only contain letter or numbers
	if !utils.IsAlphanumeric(service) && !utils.IsAlphanumeric(resource) {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, "Not a valid format")
	}

	// put in namespace dilimeter but left the resource type empty
	if resource == genericResourceType && strings.Contains(s.String(), utils.NamespaceDilimeter) {
		return utils.NewErrorf(utils.ErrorCodeInvalidArgument, "Not a valid format")
	}

	// check to see if we have paramters specififed
	if strings.Contains(s.String(), utils.ParamDilimeter) {
		_, _, err := s.SpiltParam()

		// error spliting the param not properly formated
		if err != nil {
			return utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "Not a valid format")
		}
	}	

	return nil
}

// builds a namespace from service and resource is provided. 
// Oposite of Namespace.Split.
// example: app, file -> app/file 
func BuildNamespace(service string, resource string) NameSpaceName {
	namespace := strings.Join([]string{service, resource}, utils.NamespaceDilimeter)

	return NameSpaceName(namespace)
}

// splits namespace back into two parts service and resource. 
// Oposite of Namespace.Build
// example: app/resource -> app, resource
func (s NameSpaceName) Split() (string, string) {

	// ns = ParseNamespaceAliasIfRequired(ns)

	parts := strings.Split(s.String(), utils.NamespaceDilimeter)

	// theres no namespace delimeter provided
	if len(parts) < 2 {
		return parts[0], genericResourceType
	}

	// strip out param
	if strings.Contains(parts[1], utils.ParamDilimeter) {
		pts :=  strings.Split(parts[1], utils.ParamDilimeter)
		
		return parts[0], pts[0]
	}

	return parts[0], parts[1]
}

// splits back into namespace and resource param. 
// Oposite of JoinParam
// example: app/resource:uuid -> app/resource, uuid
func (s NameSpaceName) SpiltParam() (string, string, error) {

	namespaceParts := strings.Split(s.String(), utils.NamespaceDilimeter)

	if len(namespaceParts) != 2 {
		return "", "", utils.NewErrorf(utils.ErrorCodeInvalidArgument, "bad namespace, format should namespace:uuid")
	}

	return namespaceParts[0], namespaceParts[1], nil
}

// joins namespace with resource param.
// Oposite of SpiltParam
// example: app/resource, uuid -> app/resource:uuid
func (s NameSpaceName) JoinParam(param string) string {

	param = strings.Join([]string{ s.String(), param }, utils.ParamDilimeter)

	return  param
}

// build the namespace key associated with a relation
// example: app/resource, permission -> app.resource.permission
func (s NameSpaceName) BuildKey(permission string) string {
	service, resource := s.Split()

	key := strings.Join([]string{ service, resource, permission }, utils.KeyDilimeter)

	return key
}

// builds the relation name associated with this namespace. 
// relation can be other relations, roles, permission, platforms, etc... 
// example: app/resource, owner -> app_file_owner
func (s NameSpaceName) BuildRelation(verb string) string {
	service, resource := s.Split()
	
	relation := strings.Join([]string{service, resource, verb }, utils.RelationDilimeter)
	
	return relation
}

