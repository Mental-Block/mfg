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
		Oauth    bool   `example:"false" doc:"set to true if requested to use oauth"`
	}
}

type RegisterResponse struct {
	refreshToken http.Cookie `header:"refreshtoken"`
}

func (s *ServiceInject) Register(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "register-account",
		Summary:       "register account",
		Path:          "/register/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {

		token, err := s.authService.Register(
			ctx,
			req.Body.Email,
			req.Body.Username,
			req.Body.Password,
			req.Body.Oauth,
		)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &RegisterResponse{
			refreshToken: http.Cookie{
				Name:     "bob",
				Value:    *token,
				SameSite: http.SameSiteLaxMode,
				HttpOnly: true,
			},
		}

		return resp, nil
	})
}
