package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type DeleteUserRequest struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type DeleteUserResponse struct {
	Body struct {
		Id string `json:"id" example:"1" doc:"unique identifier"`
	}
}

func (s *UserHandler) delete(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "delete-user",
		Summary:       "Delete user",
		Path:          "/{user-id}/",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
		if req.Id == "" {
			return nil, huma.Error400BadRequest("user-id can't be empty")
		}

		id, err := s.userService.Remove(ctx, req.Id)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		resp := &DeleteUserResponse{}
		
		resp.Body.Id = id

		return resp, nil
	})
}
