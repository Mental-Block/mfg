package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
)

type UserProfileRequest struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type UserProfileResponse struct {
	Body UserProfile
}

func (h *ServiceInject) getUserProfile(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "get-user",
		Summary:       "get user",
		Path:          "/{user-id}/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, input *UserProfileRequest) (*UserProfileResponse, error) {

		if input.Id == "" {
			return nil, huma.Error400BadRequest("user-id can't be empty")
		}

		id, err := strconv.Atoi(input.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("user-id has to be a number")
		}

		data, err := h.userService.GetProfile(ctx, id)

		if err != nil {
			return nil, huma.Error500InternalServerError("user store error")
		}

		resp := &UserProfileResponse{
			Body: struct {
				Id       int    `json:"id" example:"1" doc:"unique identifier"`
				Username string `json:"username" example:"bob" minLength:"1" maxLength:"30" doc:"bob"`
			}{
				Id:       int(data.Id),
				Username: string(data.Username),
			},
		}

		return resp, nil
	})
}
