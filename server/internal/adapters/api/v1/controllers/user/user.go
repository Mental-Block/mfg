package user

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/ports"
)

type UserProfile struct {
	Id       int    `json:"id" example:"1" doc:"unique identifier"`
	Username string `json:"username" example:"bob" minLength:"1" maxLength:"30" doc:"bob"`
}

type ServiceInject struct {
	userService  ports.UserService
	parrentGroup *huma.Group
}

func NewServiceInject(userService ports.UserService, group *huma.Group) *ServiceInject {
	return &ServiceInject{
		userService:  userService,
		parrentGroup: group,
	}
}

func (s *ServiceInject) Routes() {
	usersGrp := huma.NewGroup(s.parrentGroup, "/users")

	s.getUserProfile(usersGrp)
	s.getProfiles(usersGrp)
	s.deleteUser(usersGrp)
	s.isUserTaken(usersGrp)
}
