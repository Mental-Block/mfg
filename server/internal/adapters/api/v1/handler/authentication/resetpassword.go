package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type ResetPasswordRequest struct {
	Body struct {
		Email string `json:"email" example:"bob@gmail.com" doc:"unique email to each account"`
	}
}

type ResetPasswordResponse struct {
	Body struct {
		ok bool
	}
}

func (s *AuthHandler) ResetPassword(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "reset-password",
		Summary:       "Reset password",
		Description:   "resets user password for account",
		Path:          "/reset/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	
		err := s.authService.ResetPassword(ctx, req.Body.Email)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		resp := &ResetPasswordResponse{}
		resp.Body.ok = true

		return resp, nil
	})
}
