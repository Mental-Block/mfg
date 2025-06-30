package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
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

		isTaken, err := h.authService.IsEmailTaken(ctx, req.Body.Email)
		
		if (err != nil) {
			return nil, util.HumaError(err)
		}

		resp := &IsTakenResponse{
			Body: isTaken,
		}

		return resp, nil
	})
}
