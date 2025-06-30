package prefrence

import (
	"database/sql"

	"github.com/server/pkg/utils"
)

type PreferenceModel struct {
	PrefrenceId  	string    			`db:"prefrence_id"`
	ResourceId   	string    			`db:"resource_id"`
	Name         	string    			`db:"name"`
	Value        	string    			`db:"value"`
	ResourceType 	string    			`db:"resource_type"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 	`db:"created_by"`
	CreatedDt 		utils.CreatedDT 	`db:"created_dt"`
}