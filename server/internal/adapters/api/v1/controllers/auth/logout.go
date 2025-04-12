package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/env"
)

type LogoutRequest struct{}

type LogoutResponse struct {
	refreshToken http.Cookie `header:"refreshtoken"`
}

func (s *ServiceInject) Logout(api huma.API) {

	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "logout-account",
		Summary:       "logout account",
		Description:   "removes refresh token and auth token",
		Path:          "/logout/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"refreshToken": {"scope1"}},
			{"authToken": {"scope2"}},
		},
	}, func(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {

		cfg := env.Env()

		cookie := http.Cookie{
			Name:    "refresh token",
			Expires: time.Now().AddDate(0, 0, 30),
		}

		if env.Enviroment[cfg.ENV] == "development" {
			cookie.HttpOnly = false
			cookie.SameSite = http.SameSiteNoneMode
		}

		if env.Enviroment[cfg.ENV] == "production" || env.Enviroment[cfg.ENV] == "test" {
			cookie.HttpOnly = true
			cookie.SameSite = http.SameSiteStrictMode
		}

		resp := &LogoutResponse{
			refreshToken: cookie,
		}

		return resp, nil
	})
}
