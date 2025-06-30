package domain

import (
	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type Role struct {
	Id          	string        
	OrgId       	string       
	Name        	string     
	Title       	string
	Permissions 	[]string        		
	Scopes      	[]string
	Active       	bool      
	Metadata    	metadata.Metadata   
}

func (r Role)Transfrom() (*SanitizedRole, error) {

	uuid, err := utils.ConvertStringToUUID(r.Id)

	if err != nil {
		return nil, err
	}

	orguuid, err := utils.ConvertStringToUUID(r.Id)

	if err != nil {
		return nil, err
	}

	// TODO ADD SANITIZATION FOR USER ONCE I KNOW MORE OF WHAT IM DOING LOL>>>>>>>>
	roleName := RoleName(r.Name)

	title := r.Title

	active := State(r.Active)

	namespaces := make([]namespace.NameSpaceName, len(r.Scopes)) 

	for i, v := range namespaces {
		namespaces[i] = namespace.NameSpaceName(v)
	}

	permissions := make([]permission.PermissionName, len(r.Permissions))

	for i, v := range permissions {
		permissions[i] = permission.PermissionName(v) 
	}

	return &SanitizedRole{
		Id: uuid,
		OrgId: orguuid,
		Name: roleName,
		Title: title,
		Active: active,
		Scopes: namespaces,
		Permissions: permissions,
		Metadata: r.Metadata,
	}, nil
}

type SanitizedRole struct {
	Id          	utils.UUID         
	OrgId       	utils.UUID       
	Name        	RoleName  
	Title       	string   
	Active       	State
	Scopes      	[]namespace.NameSpaceName   
	Permissions 	[]permission.PermissionName        		
	Metadata    	metadata.Metadata 
}
	
func (r SanitizedRole) Transfrom() (*Role) {

	scopes := make([]string, len(r.Scopes)) 

	for i, v := range r.Scopes {
		scopes[i] = v.String()
	}

	permissions := make([]string, len(r.Permissions))

	for i, v := range r.Permissions {
		permissions[i] = v.String()
	}

	return &Role{
		Id: r.Id.String(),          	        
		OrgId: r.OrgId.String(),       	       
		Name: r.Name.String(),        	     
		Title: r.Title,        	
		Permissions: permissions, 	        		
		Scopes: scopes,      	
		Active: r.Active.Bool(),       	      
		Metadata: r.Metadata,
	}
}
