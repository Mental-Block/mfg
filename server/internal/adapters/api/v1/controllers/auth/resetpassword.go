package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type ResetPasswordRequest struct {
	Body struct {
		id          int    `example:"123" doc:"unique identifier"`
		OldPassword string `example:"MyOldPassword123!" minLength:"8" maxLength:"64" doc:"accounts current password. passwords are hased on client side and server side"`
		Password    string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"accounts new password. passwords are hased on client side and server side"`
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
		//	s.authService.ResetPassword(ctx, req.Body.)

		resp := &ResetPasswordResponse{
			Body: true,
		}

		return resp, nil
	})
}
