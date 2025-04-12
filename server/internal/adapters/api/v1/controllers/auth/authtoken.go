package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type AuthTokenRequest struct {
}

type AuthTokenResponse struct {
}

func (s *ServiceInject) AuthToken(api huma.API) {

	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "auth-token",
		Summary:       "auth permissions token",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
		Path:          "/permission/",
	}, func(ctx context.Context, i *AuthTokenRequest) (*AuthTokenResponse, error) {
		return nil, nil
	})
}
