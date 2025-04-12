package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteUserRequest struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type DeleteUserResponse struct {
	Body bool
}

func (s *ServiceInject) deleteUser(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "delete-user",
		Summary:       "delete user",
		Path:          "/{user-id}/",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"defaultAuth": {"accountHolder"}},
		},
	}, func(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
		if req.Id == "" {
			return nil, huma.Error400BadRequest("user-id can't be empty")
		}

		userId, err := strconv.Atoi(req.Id)

		if err != nil {
			return nil, huma.Error400BadRequest("user-id has to be a number")
		}

		_, err = s.userService.RemoveUser(ctx, userId)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &DeleteUserResponse{
			Body: true,
		}

		return resp, nil
	})
}
