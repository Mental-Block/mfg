package project

import (
	"database/sql"
	"time"

	"github.com/server/pkg/utils"
)


type ProjectModel struct {
	ProjectId       string         		`db:"project_id"`
	OrgId     		string         		`db:"org_id"`
	Name      		string         		`db:"name"`
	Title     		sql.NullString 		`db:"title"`
	Metadata  		[]byte         		`db:"metadata"`
	State     		sql.NullString 		`db:"state"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		utils.CreatedBy 	`db:"created_by"`
	CreatedDt 		utils.CreatedDT 	`db:"created_dt"`
}

type ProjectUsersModel struct {
	UserId          sql.NullString  `db:"user_id"`
	UserName        sql.NullString `db:"name"`
	UserEmail       sql.NullString `db:"email"`
	UserTitle       sql.NullString `db:"title"`
	UserAvatar      sql.NullString `db:"avatar"`
	UserState       sql.NullString `db:"state"`
	RoleNames       sql.NullString `db:"role_names"`
	RoleTitles      sql.NullString `db:"role_titles"`
	RoleIds         sql.NullString `db:"role_ids"`
	ProjectId       sql.NullString `db:"project_id"`
	ProjectJoinedDt sql.NullTime   `db:"project_joined_dt"`
}

type ProjectAggregatedUserModel struct {
	Id              string     `rql:"name=id,type=string"`
	Name            string     `rql:"name=name,type=string"`
	Email           string     `rql:"name=email,type=string"`
	Title           string     `rql:"name=title,type=string"`
	State           string 	   `rql:"name=state,type=string"`
	Avatar          string
	RoleNames       []string   `rql:"name=role_names,type=string"`
	RoleTitles      []string   `rql:"name=role_titles,type=string"`
	RoleIds         []string   `rql:"name=role_ids,type=string"`
	ProjectId       string     `rql:"name=project_id,type=string"`
	ProjectJoinedDT time.Time  `rql:"name=project_joined_dt,type=datetime"`
}