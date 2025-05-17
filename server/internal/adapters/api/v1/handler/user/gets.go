package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
)

type UsersResponse struct {
	Body struct {
		Users []dto.User 
	}
}

func (h *UserHandler) getUsers(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "get-users",
		Summary:       "Get users",
		Path:          "/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, _ *struct{}) (*UsersResponse, error) {

		data, err := h.userService.GetUsers(ctx)

		if err != nil  {
			return nil, util.HumaError(err)
		}
		
		users := make([]dto.User, len(data))

		for i, v := range data {
			users[i].Id = int(v.Id)
			users[i].Username = string(v.Username)
		}

		resp := &UsersResponse{}
		resp.Body.Users = users

		return resp, nil
	})

}
