package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type VerifyRequest struct {
	Token string `json:"token" path:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" doc:"email verification token"`
}

type VerifyResponse struct {
	Body struct {
		ok bool
	}
}

func (s *AuthHandler) Verify(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "verify",
		Summary:       "Verify account",
		Path:          "/verify/{token}",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
		
	}, func(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {

		// _, err := s.authService.VerifyEmailToken(ctx, req.Token)

		// if err != nil  {
		// 	return nil, util.HumaError(err)
		// }

		// resp := &VerifyResponse{}
		// resp.Body.ok = true

		return nil, nil
	})
}
