package transform

import (
	role "github.com/server/internal/core/role/domain"
	"github.com/server/pkg/metadata"
)

type Policy struct {
	PolicyId        string
	Role          	role.Role
	RoleId        	string
	ResourceId    	string
	ResourceType  	string
	PrincipalId   	string    			
	PrincipalType 	string    			
	Metadata      	metadata.Metadata    			
}

// func DBModelToNamePolicy(model PolicyModel) (*Policy, error) {
// 	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
// 	if (err != nil) {
// 		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
// 	}

// 	roleData, err := metadata.UnMarshallMetadata(model.Role.Metadata)
	
// 	if (err != nil) {
// 		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
// 	}

// 	rolePermissionData, err := metadata.UnMarshallArray(model.Role.Permissions)
	
// 	if (err != nil) {
// 		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledPermissions")
// 	}	
	
// 	role := role.Role{
// 		RoleId:         model.RoleId,  
// 		OrgId:       	model.Role.OrgId,
// 		Name:        	model.Role.Name,
// 		Title:       	model.Role.Title.String,
// 		Permissions: 	rolePermissionData,
// 		Scopes:      	model.Role.Scopes,
// 		State:       	model.Role.State,      
// 		Metadata:    	roleData,  
// 	}

// 	return &Policy{
// 		PolicyId: 		model.PolicyId,
// 		Role: 	  		role,
// 		RoleId:        	model.RoleId,    
// 		ResourceId:    	model.ResourceId,
// 		ResourceType:   model.ResourceType,
// 		PrincipalId:    model.PrincipalId,
// 		PrincipalType: 	model.PrincipalType,
// 		Metadata:       data,
// 	}, nil
// }
