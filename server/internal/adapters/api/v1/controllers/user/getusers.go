package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UserProfilesRequest struct {
	Id string `json:"id" path:"user-id" example:"1" doc:"unique identifier"`
}

type UserProfilesResponse struct {
	Body struct {
		Users []UserProfile
	}
}

func (h *ServiceInject) getProfiles(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "get-users",
		Summary:       "get users",
		Path:          "/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, _ *UserProfilesRequest) (*UserProfilesResponse, error) {

		data, err := h.userService.GetProfiles(ctx)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		users := make([]UserProfile, len(data))

		for i, v := range users {
			users[i].Id = v.Id
			users[i].Username = v.Username
		}

		resp := &UserProfilesResponse{}
		resp.Body.Users = users

		return resp, nil
	})

}
