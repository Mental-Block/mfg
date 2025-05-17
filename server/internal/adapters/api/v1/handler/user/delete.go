package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type DeleteUserRequest struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type DeleteUserResponse struct {
	Body struct {
		Id int `json:"id" example:"1" doc:"unique identifier"`
	}
}

func (s *UserHandler) deleteUser(api huma.API) {
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

		userId, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("user-id has to be a number")
		}

		_, err = s.userService.Remove(ctx, userId)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		resp := &DeleteUserResponse{}
		resp.Body.Id = userId

		return resp, nil
	})
}
