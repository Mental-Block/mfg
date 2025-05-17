package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type UserRequest struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type UserResponse struct {
	Body *dto.User
}

func (h *UserHandler) getUser(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "get-user",
		Summary:       "Get user",
		Path:          "/{user-id}/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, input *UserRequest) (*UserResponse, error) {

		if input.Id == "" {
			return nil, huma.Error400BadRequest("user-id can't be empty")
		}

		id, err := strconv.Atoi(input.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("user-id has to be a number")
		}

		data, err := h.userService.Get(ctx, id)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		resp := &UserResponse{
			Body: &dto.User{
				Id:       int(data.Id),
				Username: string(data.Username),
			},
		}

		return resp, nil
	})
}
