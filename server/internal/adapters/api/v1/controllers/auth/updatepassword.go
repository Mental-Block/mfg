package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdatePasswordRequest struct {
	Body struct {
		Password string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"account login password"`
	}
}

type UpdatePasswordResponse struct {
	Body bool
}

func (s *ServiceInject) UpdatePassword(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "update-password",
		Summary:       "update password",
		Path:          "/rests/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
		// Security: []map[string][]string{
		// 	{"defaultAuth": {"accountHolder"}},
		// },
	}, func(ctx context.Context, req *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {

		err := s.authService.UpdatePassword(ctx, "dsadsa", req.Body.Password)

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &UpdatePasswordResponse{
			Body: true,
		}

		return resp, nil
	})
}
