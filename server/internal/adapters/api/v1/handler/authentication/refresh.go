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

type RefreshRequest struct{
	Cookie http.Cookie `cookie:"mfg-refresh-token"`
}

type RefreshResponse struct {
	SetCookie []*http.Cookie `header:"Set-Cookie"`
	Body *dto.Authorization
}

func (s *AuthHandler) Refresh(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "refresh-token",
		Summary:       "Refresh token",
		Description:   "Gets a new auth token, removes the old auth token, refreshes token",
		Path:          "/refresh/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error) {

		refreshToken, authToken, auth, err := s.authService.Permission(ctx, req.Cookie.Value)

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
		refreshCookie.Value = *refreshToken;
		refreshCookie.Expires = time.Now().Add(domain.RefreshTokenDuration)

		authCookie.Name = domain.AuthTokenName;
		authCookie.Value = *authToken;
		authCookie.Expires = time.Now().Add(domain.AuthTokenDuration)

		resp := &RefreshResponse{}

		resp.Body = &dto.Authorization{
			Id: int(auth.Id),
			Username: string(auth.Username),
			Roles: auth.Roles,
		}

		resp.SetCookie = []*http.Cookie{
			&refreshCookie,	
			&authCookie,
		}

		return resp, nil
	})
}
