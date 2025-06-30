package inventation

import "time"

type InvitationModel struct {
	InvitationId    string 	  `db:"invitation_id"`
	UserId    		string    `db:"user_id"`
	OrgId    		string    `db:"org_id"`
	Metadata  		[]byte    `db:"metadata"`
	CreatedDT 		time.Time `db:"created_at"`
	CreatedBy 		string    `db:"created_by"`
	ExpiresDT 		time.Time `db:"expires_at"`
}