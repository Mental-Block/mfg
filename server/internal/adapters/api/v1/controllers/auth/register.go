package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type RegisterRequest struct {
	Body struct {
		Username string `example:"bob" minLength:"1" maxLength:"30" doc:"bob"`
		Password string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"accounts new password"`
		Email    string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
	}
}

type RegisterResponse struct {}

func (s *ServiceInject) Register(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "register-account",
		Summary:       "register account",
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

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &RegisterResponse{}

		return resp, nil
	})
}
