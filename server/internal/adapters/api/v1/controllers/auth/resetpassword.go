package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type ResetPasswordRequest struct {
	Body struct {
		Email string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
	}
}

type ResetPasswordResponse struct {
	Body bool
}

func (s *ServiceInject) ResetPassword(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "reset-password",
		Summary:       "reset password",
		Description:   "resets user password for account",
		Path:          "/reset/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	
		err := s.authService.ResetPassword(ctx, req.Body.Email)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &ResetPasswordResponse{
			Body: true,
		}

		return resp, nil
	})
}
