package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/domain/entity"
	"github.com/server/internal/core/services/auth"
)

type DeleteAccountInput struct {
	Body struct {
		Id int `example:"1" doc:"unique identifier"`
	}
}

type DeleteAccountOutput struct {
	Body struct {
		Id entity.Id
	}
}

func (h *Handler) registerDeleteAccount(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "delete-account",
		Summary:       "delete account",
		Path:          "/",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *DeleteAccountInput) (*DeleteAccountOutput, error) {

		user, err := h.authService.RemoveAccount(ctx, auth.RemoveAccountInput(req.Body))

		if err != nil {
			return nil, huma.Error500InternalServerError(err.Error())
		}

		resp := &DeleteAccountOutput{
			Body: struct{ Id entity.Id }{Id: user.Id},
		}

		return resp, nil
	})
}
