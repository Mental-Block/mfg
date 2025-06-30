package user

import (
	"database/sql"
	"time"

	"github.com/server/internal/core/user/domain"
	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type UserModel struct {
	Id 				utils.UUID			`db:"user_id"`
	Username 		domain.Username		`db:"username"`	
	Active          bool				`db:"active"`
	Title     		domain.Title 		`db:"title"`
	Avatar    		domain.Avtar 		`db:"avatar"`
	Metadata  		[]byte			    `db:"metadata"`
	DeletedBy 		sql.NullString 		`db:"deleted_by"`
	DeletedDT 		sql.NullTime 		`db:"deleted_dt"`
	UpdateBy 		sql.NullString 	 	`db:"updated_by"`
	UpdatedDT 		sql.NullTime 		`db:"updated_dt"`
	CreatedBy 		string 				`db:"created_by"`
	CreatedDt 		time.Time			`db:"created_dt"`
}

func (model UserModel) Transform()(domain.SanitizedUser, error) {
	data, err := metadata.UnMarshallMetadata(model.Metadata)
	
	if (err != nil) {
		return domain.SanitizedUser{}, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "metadata.UnMarshallMetadata")
	}

	return domain.SanitizedUser{
		Id:           	model.Id,
		Username: 		model.Username,
		Active:         model.Active,
		Title:    		model.Title,
		Avatar:    		model.Avatar,
		Metadata:  		data,
	}, nil
}


type UserProjectsModel struct {
	ProjectId        sql.NullString 	`db:"project_id"`
	OrgId            sql.NullString 	`db:"org_id"`
	UserId           sql.NullString 	`db:"user_id"`
	ProjectTitle     sql.NullString 	`db:"project_title"`
	ProjectName      sql.NullString 	`db:"project_name"`
	ProjectCreatedDt sql.NullTime   	`db:"project_created_dt"`
	UserIds          []string 			`db:"user_ids"`
	UserNames        []string 			`db:"user_names"`
	UserTitles       []string 			`db:"user_titles"`
	UserAvatars      []string 			`db:"user_avatars"`
}

type UserOrgsModel struct {
	OrgId        sql.NullString `db:"org_id"`
	OrgTitle     sql.NullString `db:"org_title"`
	OrgName      sql.NullString `db:"org_name"`
	OrgAvatar    sql.NullString `db:"org_avatar"`
	ProjectCount sql.NullInt64  `db:"project_count"`
	RoleNames    []string 		`db:"role_names"`
	RoleTitles   []string 		`db:"role_titles"`
	RoleIds      []string 		`db:"role_ids"`
	OrgJoinedOn  sql.NullTime   `db:"org_joined_on"`
	UserID       sql.NullString `db:"principal_id"`
}

type UserOrgsGroupModel struct {
	Name sql.NullString      `db:"name"`
	Data []UserOrgsGroupDataModel `db:"data"`
}

type UserOrgsGroupDataModel struct {
	Name  sql.NullString `db:"values"`
	Count int            `db:"count"`
}