package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
	"github.com/server/internal/core/user/domain"
)

type UpdateUserRequest struct{
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
	Body User
}

type UpdateUserResponse struct {
	Body User
}

func (s *UserHandler) update(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "update-user",
		Summary:       "Update user",
		Description:   "Updates user",
		Path:          "/{user-id}/",
		Method:        http.MethodPatch,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
		
		if req.Id == "" {
			return nil, huma.Error400BadRequest("user-id can't be empty")
		}

		user, err := s.userService.Update(ctx, domain.User(req.Body))

		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &UpdateUserResponse{}

		resp.Body = User(*user)

		return resp, nil
	})
}