package resource

import (
	"database/sql"

	"github.com/server/internal/core/resource/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)
	
type ResourceModel struct {
	Id            	string         				`db:"resource_id"`
	URN           	string         				`db:"urn"`
	Name          	string         				`db:"name"`
	Title         	string         				`db:"title"`
	ProjectId     	string         				`db:"project_id"`
	NamespaceId   	string         				`db:"namespace_name"`
	PrincipalId   	sql.NullString 				`db:"principal_id"`
	PrincipalType 	sql.NullString 				`db:"principal_type"`
	Metadata      	[]byte         				`db:"metadata"`
	DeletedBy 		sql.NullString 				`db:"deleted_by"`
	DeletedDT 		sql.NullTime 				`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 			`db:"updated_by"`
	UpdatedDT 		sql.NullTime 				`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 			`db:"created_by"`
	CreatedDt 		utils.CreatedDT 			`db:"created_dt"`
}

func (r ResourceModel) Transform()(domain.SanitizedResource, error) {
	data, err := metadata.UnMarshallMetadata(r.Metadata)
	
	if (err != nil) {
		return domain.SanitizedResource{}, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	return domain.SanitizedResource{
		Id:     r.Id,
		PrincipalId:   r.PrincipalId.String,
		PrincipalType: r.PrincipalType.String,
		URN:           r.URN,
		Name:          r.Name,
		Title:         r.Title,
		ProjectId:     r.ProjectId,
		NamespaceId:   r.NamespaceId,
		Metadata:      data,
	}, nil
}

