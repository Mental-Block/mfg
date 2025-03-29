package user

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/server/internal/core/domain/entity"
	domain "github.com/server/internal/core/domain/user"
	"github.com/server/internal/core/services/user"
)

type UserAPI struct {
	Username domain.Username `example:"bob" minLength:"1" maxLength:"30" doc:"my username is 'bob'"`
}

type UserEntityAPI struct {
	Id        entity.Id        `example:"123" doc:"unique identifier"`
	Username  domain.Username  `example:"bob" minLength:"1" maxLength:"30" doc:"bob"`
	CreatedBy entity.CreatedBy `example:"bob" doc:"created by bob"`
	CreatedDT entity.CreatedDT `example:"2025-03-19T15:24:00.869142-04:00" doc:"created using UTC"`
	UpdatedBy entity.UpdatedBy `example:"bob" doc:"updated by bob"`
	UpdatedDT entity.UpdatedDT `example:"2025-03-19T15:24:00.869142-04:00" doc:"created using UTC"`
}

type Handler struct {
	userService user.UserService
}

func NewHandler(userService user.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) Routes(group *huma.Group) {
	api := huma.NewGroup(group, "/users")

	h.registerGetUser(api)
	h.registerGetUsers(api)
}
