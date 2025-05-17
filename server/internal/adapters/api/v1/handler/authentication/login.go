package authentication

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/dto"
	"github.com/server/internal/adapters/api/v1/util"
	"github.com/server/internal/adapters/env"
	"github.com/server/internal/core/domain"
)

type LoginRequest struct {
	Body struct {
		Email    string `example:"bob@gmail.com" doc:"unique email to each account"`
		Password string `example:"MyNewPassword123!" doc:"account login password"`
	}
}

type LoginResponse struct {
	SetCookie []*http.Cookie `header:"Set-Cookie"`
	Body *dto.Authorization
}

func (s *AuthHandler) Login(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "login-account",
		Summary:       "Login account",
		Description:   "validates user credentials, returning a refresh token",
		Path:          "/login/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
		
		rToken, aToken, auth, err := s.authService.Login(
			ctx,
			req.Body.Email,
			req.Body.Password,
		)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		cfg := env.Env()

		cookie := http.Cookie{
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

		refreshCookie := cookie
		authCookie := cookie

		refreshCookie.Name = domain.RefreshTokenName;
		refreshCookie.Value = *rToken;
		refreshCookie.Expires = time.Now().Add(domain.RefreshTokenDuration)

		authCookie.Name = domain.AuthTokenName;
		authCookie.Value = *aToken;
		authCookie.Expires = time.Now().Add(domain.AuthTokenDuration)

		resp := &LoginResponse{}

		resp.SetCookie = []*http.Cookie{
			&refreshCookie,	
			&authCookie,
		}

		resp.Body = &dto.Authorization{
			Id: int(auth.Id),
			Username: string(auth.Username),
			Roles: auth.Roles,
		}

		return resp, nil
	})
}
