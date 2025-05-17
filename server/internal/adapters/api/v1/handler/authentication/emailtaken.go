package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
	"github.com/server/internal/core/domain"
)

type IsTakenRequest struct {
	Body struct {
		Email string `json:"email" example:"bob@gmail.com" doc:"unique email to each account"`
	}
}

type IsTakenResponse struct {
	Body bool `example:"true" doc:"check if user exist"`
}

func (h *AuthHandler) EmailTaken(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "email-taken",
		Summary:       "Email taken",
		Path:          "/email-taken/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *IsTakenRequest) (*IsTakenResponse, error) {

		_, err := h.userService.GetUserByEmail(ctx, req.Body.Email)
		
		resp := &IsTakenResponse{}
		
		if (err != nil) {
			if (err.Error() == domain.ErrUserNotFound.Error()) {
				resp.Body = false
				return resp, nil
			}

			return nil, util.HumaError(err)
		}

		resp.Body = true
		return resp, nil
	})
}
