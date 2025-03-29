package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetUsersOuput struct {
	Body struct {
		Users []UserEntityAPI `json:"users"`
	}
}

func (h *Handler) registerGetUsers(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-users",
		Summary:       "get users",
		Path:          "/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, _ *struct{}) (*GetUsersOuput, error) {
		data, err := h.userService.GetUsers(ctx)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		userEntity := make([]UserEntityAPI, len(data))

		for i, v := range data {
			userEntity[i] = UserEntityAPI(v)
		}

		resp := &GetUsersOuput{}
		resp.Body.Users = userEntity

		return resp, nil
	})

}
