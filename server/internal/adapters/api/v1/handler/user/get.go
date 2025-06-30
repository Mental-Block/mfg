package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type UserRequest struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type UserResponse struct {
	Body *User
}

func (h *UserHandler) get(api huma.API) {
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
		
		user, err := h.userService.Get(ctx, input.Id)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		resp := &UserResponse{
			Body: &User{
				Id: user.Id,
				Username: user.Username,
				Active: user.Active,
				Title: user.Title,
				Avatar: user.Title,
				Metadata: user.Metadata,
			},
		}

		return resp, nil
	})
}
