package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type isTakenRequest struct {
	Body struct {
		Email string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
	}
}

type isTakenResponse struct {
	Body struct {
		Value bool `example:"true" doc:"check if user exist"`
	} 
}

func (h *ServiceInject) isUserTaken(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"user"},
		OperationID:   "is-user-taken",
		Summary:       "is user taken?",
		Path:          "/is-taken/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *isTakenRequest) (*isTakenResponse, error) {

		data, err := h.userService.GetUser(ctx, req.Body.Email)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &isTakenResponse{}

		if data != nil {
			resp.Body.Value = true
		} else {
			resp.Body.Value = false
		}
		
		return resp, nil
	})
}
