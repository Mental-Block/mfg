package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/domain"
)

type LogoutRequest struct{
	Cookie http.Cookie `cookie:"mfg-refresh-token"`
}

type LogoutResponse struct {
	SetCookie []*http.Cookie `header:"Set-Cookie"`
}

func (s *ServiceInject) Logout(api huma.API) {

	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "logout-account",
		Summary:       "logout account",
		Description:   "Removes refresh token and auth token. Logging out the user",
		Path:          "/logout/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{domain.AuthTokenName: {"scope1"}},
		},
	}, func(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {

		err := s.authService.Logout(ctx, req.Cookie.Value)

		if (err != nil) {
			huma.Error400BadRequest("%v", err)
		}

		resp := &LogoutResponse{}

		resp.SetCookie = []*http.Cookie{
			{
				Name: domain.RefreshTokenName,
				Value:   "",
				Expires:  time.Unix(0, 0),
				MaxAge:   -1,
				Path: "/",
			},
			{
				Name: domain.AuthTokenName,
				Value:   "",
				Expires:  time.Unix(0, 0),
				MaxAge:   -1,
				Path: "/",
			},
		}

		return resp, nil
	})
}
