package namespace_store

import (
	"database/sql"

	"github.com/server/internal/core/namespace/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type NamespaceModel struct {
	Id        		utils.UUID   				`db:"namespace_id"`
	Name      		domain.NameSpaceName   		`db:"name"`
	Metadata  		[]byte       				`db:"metadata"`
	DeletedBy 		sql.NullString 				`db:"deleted_by"`
	DeletedDT 		sql.NullTime 				`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 			`db:"updated_by"`
	UpdatedDT 		sql.NullTime 				`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 			`db:"created_by"`
	CreatedDt 		utils.CreatedDT 			`db:"created_dt"`
}

func (model NamespaceModel) Transform() (domain.SanitizedNamespace, error) {
	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return domain.SanitizedNamespace{}, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "unMarshalledMetadata")
	}

	return domain.SanitizedNamespace{
		Id:    			model.Id,
		Name:      		model.Name,
		Metadata:  		data,
	}, nil
}
