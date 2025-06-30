package metaschema

import (
	"database/sql"

	"github.com/server/pkg/utils"
)

type MetaSchemaModel struct {
	Id        		string    			`db:"id"`
	Name      		string    			`db:"name"`
	Schema    		string    			`db:"schema"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 	`db:"created_by"`
	CreatedDt 		utils.CreatedDT 	`db:"created_dt"`
}

