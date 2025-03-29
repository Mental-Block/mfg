package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	API "github.com/server/internal/adapters/web/user"
	"github.com/server/internal/core/services/auth"
)

type CreateAccountInput struct {
	Body struct {
		Username string `example:"bob" minLength:"1" maxLength:"30" doc:"bob"`
		Password string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"accounts new password. passwords are hased on client side and server side"`
		Email    string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
	}
}

type CreateAccountOutput struct {
	Body API.UserEntityAPI
}

func (h *Handler) registerPostAccount(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-account",
		Summary:       "create account",
		Path:          "/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *CreateAccountInput) (*CreateAccountOutput, error) {

		account, err := h.authService.CreateAccount(ctx, auth.CreateAccountInput(req.Body))

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &CreateAccountOutput{
			Body: API.UserEntityAPI(*account),
		}

		return resp, nil
	})
}
