package role

import (
	"database/sql"
	"time"

	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"
	role "github.com/server/internal/core/role/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type RoleModel struct {
	Id          	utils.UUID         			`db:"role_id"`
	OrgId       	utils.UUID       			`db:"organization_id"`
	Name        	role.RoleName        		`db:"name"`
	Title       	sql.NullString 				`db:"title"`
	Active       	role.State        			`db:"active"`	
	Permissions 	[]byte						`db:"permissions"`
	Scopes      	[]namespace.NameSpaceName	`db:"scopes"`
	Metadata    	[]byte         				`db:"metadata"`
	DeletedBy 		sql.NullString 				`db:"deleted_by"`
	DeletedDT 		sql.NullTime 				`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 			`db:"updated_by"`
	UpdatedDT 		sql.NullTime 				`db:"updated_dt"`
	CreatedBy 	  	string 			    		`db:"created_by"`
	CreatedDt 	  	time.Time					`db:"created_dt"`
}

func (r RoleModel) Transform() (*role.SanitizedRole, error) {

	unMarshallMetadataData, err := metadata.UnMarshallMetadata(r.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "metadata.UnMarshallMetadata")
	}

	unMarshallPermissions, err := metadata.UnMarshallArray[permission.PermissionName](r.Permissions)

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "metadata.UnMarshallArray")
	}

	return &role.SanitizedRole{
		Id: r.Id, 
		OrgId: r.OrgId,             
		Name: r.Name,           
		Title: r.Title.String,     		
		Scopes: r.Scopes,  
		Active: r.Active, 
		Permissions: unMarshallPermissions,           
		Metadata: unMarshallMetadataData,
	}, nil
}
