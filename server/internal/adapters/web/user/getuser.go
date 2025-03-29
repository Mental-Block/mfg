package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/services/user"
)

type GetUserInput struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type GetUserOuput struct {
	Body UserEntityAPI
}

func (h *Handler) registerGetUser(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-user",
		Summary:       "get user",
		Path:          "/{user-id}",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, input *GetUserInput) (*GetUserOuput, error) {

		if input.Id == "" {
			return nil, huma.Error400BadRequest("user-id can't be empty")
		}

		id, err := strconv.Atoi(input.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("user-id has to be a number")
		}

		data, err := h.userService.GetUser(ctx, user.GetUserInput(id))

		if err != nil {
			return nil, huma.Error500InternalServerError("user store error")
		}

		resp := &GetUserOuput{}
		resp.Body = UserEntityAPI(*data)

		return resp, nil
	})

}
