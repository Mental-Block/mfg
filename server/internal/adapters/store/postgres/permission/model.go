package permission

import (
	"database/sql"

	namespace "github.com/server/internal/core/namespace/domain"
	permission "github.com/server/internal/core/permission/domain"

	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type PermissionModel struct {
	Id    			utils.UUID    				`db:"permission_id"`
	Name        	permission.PermissionName   `db:"name"`
	Slug        	permission.PermissionSlug   `db:"slug"`
	NamespaceId		namespace.NameSpaceName     `db:"namespace_name"`
	Metadata    	[]byte    					`db:"metadata"`
	DeletedBy 		sql.NullString 				`db:"deleted_by"`
	DeletedDT 		sql.NullTime 				`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 			`db:"updated_by"`
	UpdatedDT 		sql.NullTime 				`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 			`db:"created_by"`
	CreatedDt 		utils.CreatedDT 			`db:"created_dt"`
}

func (p PermissionModel) Transform() (*permission.SanitizedPermission, error) {
	data, err := metadata.UnMarshallMetadata(p.Metadata)
	
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	return &permission.SanitizedPermission{
		Id: p.Id,
		Name: p.Name,
		Slug: p.Slug, 	
		NamespaceId: p.NamespaceId,
		Metadata: data,
	}, nil
}
