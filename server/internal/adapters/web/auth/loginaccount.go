package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	API "github.com/server/internal/adapters/web/user"
	"github.com/server/internal/core/services/auth"
)

type LoginAccountInput struct {
	Body struct {
		Email    string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
		Password string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"accounts new password. passwords are hased on client side and server side"`
	}
}

type LoginAccountOutput struct {
	Body API.UserEntityAPI
}

func (h *Handler) registerLoginAccount(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "login-account",
		Summary:       "login account",
		Description:   "validates user creds, grabing auth token",
		Path:          "/login/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, userInput *LoginAccountInput) (*LoginAccountOutput, error) {

		account, err := h.authService.LoginAccount(ctx, auth.LoginAccountInput(userInput.Body))

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &LoginAccountOutput{
			Body: API.UserEntityAPI(*account),
		}

		return resp, nil
	})

}
