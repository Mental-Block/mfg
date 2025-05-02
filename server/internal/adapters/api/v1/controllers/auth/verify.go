package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type VerifyRequest struct {
	Token string `json:"token" path:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" doc:"email verification token"`
}

type VerifyResponse struct {
	Body bool
}

func (s *ServiceInject) Verify(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "verify",
		Summary:       "verify account",
		Path:          "/verify/{token}",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
		
	}, func(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {

		if req.Token == "" {
			return nil, huma.Error400BadRequest("token can't be empty")
		}

		_, err := s.authService.Verify(ctx, req.Token)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &VerifyResponse{
			Body: true,
		}

		return resp, nil
	})
}
