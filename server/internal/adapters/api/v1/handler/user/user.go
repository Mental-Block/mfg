package user

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/user"
	"github.com/server/pkg/metadata"
)

type User struct {
	Id       		string    			`json:"id" example:"1" doc:"unique identifier"`
	Username 		string 				`json:"username" example:"bob" doc:"bob"`
	Active          bool				`json:"active" example:"true" doc:"user is active or not."`
	Title     		string				`json:"title" example:"Mrs" doc:"users title"`
	Avatar    		string 				`json:"avatar" example:"file://endpoint/to/blob/storage"`
	Metadata  		metadata.Metadata   `json:"metadata" example:"{}" `
}

type UserHandler struct {
	userService  	user.UserService
}

func NewUserHandler(
	service user.UserService,
) *UserHandler {
	return &UserHandler{
		userService:  service,
	}
}

func (s *UserHandler) Routes(parrentGrp *huma.Group) {
	usersGrp := huma.NewGroup(parrentGrp, "/users")

	s.gets(usersGrp)
	s.get(usersGrp)
	s.delete(usersGrp)
	s.update(usersGrp)
}