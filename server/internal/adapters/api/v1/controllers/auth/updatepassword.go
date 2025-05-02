package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdatePasswordRequest struct {
	Token string `json:"token" path:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" doc:"reset password token"`
	Body struct {
		Password string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"account login password"`
	}
}

type UpdatePasswordResponse struct {
	Body bool
}

func (s *ServiceInject) UpdatePassword(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "update-password",
		Summary:       "update password",
		Path:          "/reset/{token}",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {

		err := s.authService.UpdatePassword(ctx, req.Token, req.Body.Password)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &UpdatePasswordResponse{
			Body: true,
		}

		return resp, nil
	})
}
