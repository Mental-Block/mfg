package user

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/ports"
)

type UserHandler struct {
	userService  	ports.UserService
}

func NewUserHandler(
	service ports.UserService,
) *UserHandler {
	return &UserHandler{
		userService:  service,
	}
}

func (s *UserHandler) Routes(parrentGrp *huma.Group) {
	usersGrp := huma.NewGroup(parrentGrp, "/users")

	s.getUsers(usersGrp)
	s.getUser(usersGrp)
	s.deleteUser(usersGrp)
}
