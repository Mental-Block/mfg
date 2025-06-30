package domain

import (
	"database/sql"

	"github.com/server/pkg/utils"
)

type DomainModel struct {
	DomainId        string    			`db:"domain_id"`
	OrgId     		string    			`db:"org_id"`
	Name      		string    			`db:"name"`
	Token     		string    			`db:"token"`
	State     		string    			`db:"state"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 	`db:"created_by"`
	CreatedDt 		utils.CreatedDT 	`db:"created_dt"`
}

