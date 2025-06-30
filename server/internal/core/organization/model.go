package organization

import (
	"database/sql"
	"time"

	"github.com/server/pkg/utils"
)

type OrganizationModel struct {
	OrgId  			string         		`db:"org_id"`
	Name      		string         		`db:"name"`
	Title     		sql.NullString 		`db:"title"`
	Avatar    		sql.NullString 		`db:"avatar"`
	Metadata  		[]byte         		`db:"metadata"`
	State     		sql.NullString 		`db:"state"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 	`db:"created_by"`
	CreatedDt 		utils.CreatedDT 	`db:"created_dt"`
}

type OrgUsersModel struct {
	UserId      sql.NullString `db:"user_id"`
	UserTitle   sql.NullString `db:"title"`
	UserName    sql.NullString `db:"name"`
	UserEmail   sql.NullString `db:"email"`
	UserState   sql.NullString `db:"state"`
	UserAvatar  sql.NullString `db:"avatar"`
	RoleNames   []string 	   `db:"role_names"`
	RoleTitles  []string 	   `db:"role_titles"`
	RoleIds     []string 	   `db:"role_ids"`
	OrgId       sql.NullString `db:"org_id"`
	OrgJoinedAt sql.NullTime   `db:"org_joined_at"`
}

type OrgUsersGroupModel struct {
	Name sql.NullString      `db:"name"`
	Data []OrgUsersGroupDataModel `db:"data"`
}

type OrgUsersGroupDataModel struct {
	Name  sql.NullString `db:"values"`
	Count int            `db:"count"`
}

type GroupDataModel struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type OrgAggregatedUserModel struct {
	ID          string     `rql:"name=id,type=string"`
	Name        string     `rql:"name=name,type=string"`
	Title       string     `rql:"name=title,type=string"`
	Avatar      string     `rql:"name=avatar,type=string"`
	Email       string     `rql:"name=email,type=string"`
	State       string 	   `rql:"name=state,type=string"`
	RoleNames   []string   `rql:"name=role_names,type=string"`
	RoleTitles  []string   `rql:"name=role_titles,type=string"`
	RoleIDs     []string   `rql:"name=role_ids,type=string"`
	OrgID       string     `rql:"name=org_id,type=string"`
	OrgJoinedAt time.Time  `rql:"name=org_joined_at,type=datetime"`
}

type OrgTokenModel struct {
	Amount      sql.NullInt64  `db:"token_amount"`
	Type        sql.NullString `db:"token_type"`
	Description sql.NullString `db:"token_description"`
	UserID      sql.NullString `db:"token_user_id"`
	UserTitle   sql.NullString `db:"user_title"`
	UserAvatar  sql.NullString `db:"user_avatar"`
	CreatedAt   sql.NullTime   `db:"token_created_at"`
	OrgID       sql.NullString `db:"org_id"`
}


type OrgProjectsModel struct {
	ProjectID      sql.NullString `db:"id"`
	ProjectName    sql.NullString `db:"name"`
	ProjectTitle   sql.NullString `db:"title"`
	ProjectState   sql.NullString `db:"state"`
	MemberCount    sql.NullInt64  `db:"member_count"`
	UserIDs        []string  `db:"user_ids"`
	CreatedAt      sql.NullTime   `db:"created_at"`
	OrganizationID sql.NullString `db:"org_id"`
}

type OrgProjectsGroupModel struct {
	Name sql.NullString         `db:"name"`
	Data []OrgProjectsGroupDataModel `db:"data"`
}

type OrgProjectsGroupDataModel struct {
	Name  sql.NullString `db:"values"`
	Count int            `db:"count"`
}