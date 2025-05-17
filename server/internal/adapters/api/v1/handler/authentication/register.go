package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type RegisterRequest struct {
	Body struct {
		Username string `json:"username" example:"bob" doc:"bob"`
		Password string `json:"password" example:"MyNewPassword123!" doc:"accounts new password"`
		Email    string `json:"email" example:"bob@gmail.com" doc:"unique email to each account"`
	}
}

type RegisterResponse struct {
	Body bool
}

func (s *AuthHandler) Register(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "register-account",
		Summary:       "Register account",
		Path:          "/register/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {

		err := s.authService.Register(
			ctx,
			req.Body.Email,
			req.Body.Username,
			req.Body.Password,
		)

		if err != nil  {
			return nil, util.HumaError(err)
		}
		
		resp := &RegisterResponse{}
		resp.Body = true

		return resp, nil
	})
}
