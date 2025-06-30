package audit

import "github.com/server/pkg/utils"


type AuditModel struct {
	AuditId     string 				`db:"audit_id"`
	OrgId  		string 				`db:"org_id"`
	Source 		string 				`db:"source"`
	Action 		string 				`db:"action"`
	Actor 	    []byte			    `db:"actor"`
	Target	    []byte			    `db:"target"`
	Metadata 	[]byte			    `db:"metadata"`
	CreatedBy 	utils.CreatedBy 	`db:"created_by"`
	CreatedDt 	utils.CreatedDT 	`db:"created_dt"`
}
