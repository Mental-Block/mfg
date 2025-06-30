package transform

import (
	"database/sql"

	"github.com/server/pkg/utils"
)


type PolicyModel struct {
	PolicyId        string 				`db:"policy_id"`
//	Role          	role.RoleModel				
	RoleId        	string    			`db:"role_id"`
	ResourceId    	string    			`db:"resource_id"`
	ResourceType  	string    			`db:"resource_type"`
	PrincipalId   	string    			`db:"principal_id"`
	PrincipalType 	string    			`db:"principal_type"`
	Metadata      	[]byte    			`db:"metadata"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 	`db:"created_by"`
	CreatedDt 		utils.CreatedDT 	`db:"created_dt"`
}
