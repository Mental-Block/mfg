package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type UsersResponse struct {
	Body struct {
		Users []User `json:"users"`
	}
}

func (h *UserHandler) gets(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "get-users",
		Summary:       "Get users",
		Path:          "/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, _ *struct{}) (*UsersResponse, error) {

		data, err := h.userService.List(ctx)

		if err != nil  {
			return nil, util.HumaError(err)
		}
		
		users := make([]User, len(data))

		for i, v := range data {
			users[i].Id = v.Id
			users[i].Username = v.Username
			users[i].Active = v.Active
			users[i].Avatar = v.Avatar
			users[i].Title = v.Title
			users[i].Metadata = v.Metadata
		}

		resp := &UsersResponse{}
		
		resp.Body.Users = users

		return resp, nil
	})

}
