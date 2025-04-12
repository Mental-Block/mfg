package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/env"
)

type LoginRequest struct {
	Body struct {
		Email    string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
		Password string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"account login password"`
		OAuth    bool   `example:"false" doc:"wheather account is using oauth or not"`
	}
}

type LoginResponse struct {
	refreshToken http.Cookie `header:"Set-Cookie"`
}

func (s *ServiceInject) Login(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "login-account",
		Summary:       "login account",
		Description:   "validates user credentials, returning a refresh token",
		Path:          "/login/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {

		_, err := s.authService.Login(
			ctx,
			req.Body.Email,
			req.Body.Password,
			req.Body.OAuth,
		)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		cfg := env.Env()

		cookie := http.Cookie{
			Name:    "refresh token",
			Expires: time.Now().AddDate(0, 0, 30),
		}

		if env.Enviroment[cfg.ENV] == "development" {
			cookie.HttpOnly = false
			cookie.SameSite = http.SameSiteNoneMode
			cookie.Domain = "localhost"
			cookie.Secure = false
		}

		if env.Enviroment[cfg.ENV] == "production" || env.Enviroment[cfg.ENV] == "test" {
			cookie.HttpOnly = true
			cookie.SameSite = http.SameSiteStrictMode
		}

		resp := &LoginResponse{
			refreshToken: cookie,
		}

		return resp, nil
	})
}
