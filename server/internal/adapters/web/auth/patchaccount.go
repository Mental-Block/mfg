package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	API "github.com/server/internal/adapters/web/user"
	"github.com/server/internal/core/services/auth"
)

type UpdateAccountInput struct {
	Body struct {
		Id          int    `example:"1" doc:"unique identifier"`
		Username    string `example:"bob" minLength:"1" maxLength:"30" doc:"bob"`
		Email       string `example:"bob@gmail.com" maxLength:"255" doc:"unique email to each account"`
		OldPassword string `example:"MyOldPassword123!" minLength:"8" maxLength:"64" doc:"accounts current password. passwords are hased on client side and server side"`
		Password    string `example:"MyNewPassword123!" minLength:"8" maxLength:"64" doc:"accounts new password. passwords are hased on client side and server side"`
	}
}

type UpdateAccountOutput struct {
	Body API.UserEntityAPI
}

func (h *Handler) registerPatchAccount(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "update-account",
		Summary:       "update account",
		Path:          "/",
		Method:        http.MethodPatch,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *UpdateAccountInput) (*UpdateAccountOutput, error) {

		user, err := h.authService.UpdateAccount(ctx, auth.UpdateAccountInput(req.Body))

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &UpdateAccountOutput{
			Body: API.UserEntityAPI(*user),
		}

		return resp, nil
	})
}
