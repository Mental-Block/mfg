package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func (s *ServiceInject) Refresh(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "refresh-token",
		Summary:       "refresh token",
		Description:   "Gets a new refresh token, removes the old refresh token.",
		Path:          "/refresh/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *struct{}) (*struct{}, error) {

		
			







		return nil, nil
	})
}
