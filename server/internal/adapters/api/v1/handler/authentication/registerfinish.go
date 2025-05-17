package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/util"
)

type RegisterFinishRequest struct {
	Token string `json:"token" path:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" doc:"email verification token for finishing registration"`
}

type RegisterFinishResponse struct {
	Body bool
}

func (s *AuthHandler) RegisterFinish(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "finish-register",
		Summary:       "Finish register account",
		Path:          "/finish-register/{token}",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *VerifyRequest) (*RegisterFinishResponse, error) {

		if req.Token == "" {
			return nil, huma.Error400BadRequest("token can't be empty")
		}

		err := s.authService.RegisterFinish(ctx, req.Token)

		if err != nil  {
			return nil, util.HumaError(err)
		}

		resp := &RegisterFinishResponse{
			Body: true,
		}

		return resp, nil
	})
}
