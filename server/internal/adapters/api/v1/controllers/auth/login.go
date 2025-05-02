package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/env"
	"github.com/server/internal/core/domain"
)

type LoginRequest struct {
	Body struct {
		Email    string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
		Password string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"account login password"`
	}
}

type LoginResponse struct {
	SetCookie http.Cookie `header:"Set-Cookie"`
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

		token, err := s.authService.Login(
			ctx,
			req.Body.Email,
			req.Body.Password,
		)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		cfg := env.Env()

		cookie := http.Cookie{
			Name:    domain.RefreshTokenName,
			Value:   *token,
			Expires: time.Now().Add(domain.RefreshTokenDuration),
			Path: "/",
		}

		if cfg.ENV == env.Development {
			cookie.Secure = false	
			cookie.HttpOnly = false
			cookie.SameSite = http.SameSiteLaxMode
		}

		if cfg.ENV == env.Production || cfg.ENV == env.Test {
			cookie.Domain = cfg.SMTP.Host
			cookie.HttpOnly = true
			cookie.Secure = true
			cookie.SameSite = http.SameSiteStrictMode
		}

		resp := &LoginResponse{
			SetCookie: cookie,
		}

		return resp, nil
	})
}
