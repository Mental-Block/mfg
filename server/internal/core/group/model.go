package group

import (
	"database/sql"

	"github.com/server/pkg/utils"
)


type GroupModel struct {
	GroupId        	string    			`db:"group_id"`
	Name      		string         		`db:"name"`
	Title     		sql.NullString 		`db:"title"`
	OrgId     		string         		`db:"org_id"`
	State     		sql.NullString 		`db:"state"`
	Metadata  		[]byte       		`db:"metadata"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 	`db:"created_by"`
	CreatedDt 		utils.CreatedDT 	`db:"created_dt"`
}