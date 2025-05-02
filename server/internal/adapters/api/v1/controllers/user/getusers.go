package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UserProfilesRequest struct {}

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

		for i, v := range data {
			users[i].Id = int(v.Id)
			users[i].Username = string(v.Username)
		}

		resp := &UserProfilesResponse{}
		resp.Body.Users = users

		return resp, nil
	})

}
