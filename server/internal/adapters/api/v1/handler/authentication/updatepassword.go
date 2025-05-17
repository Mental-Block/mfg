package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type UpdatePasswordRequest struct {
	Token string `json:"token" path:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" doc:"reset password token"`
	Body struct {
		Password string `json:"password" example:"MyNewPassword123!" doc:"account login password"`
	}
}

type UpdatePasswordResponse struct {
	Body struct {
		ok bool
	}
}

func (s *AuthHandler) UpdatePassword(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "update-password",
		Summary:       "Update password",
		Path:          "/reset/{token}",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {

		err := s.authService.UpdatePassword(ctx, req.Token, req.Body.Password)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		resp := &UpdatePasswordResponse{}
		resp.Body.ok = true

		return resp, nil
	})
}
