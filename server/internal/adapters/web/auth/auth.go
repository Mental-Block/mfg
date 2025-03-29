package auth

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/services/auth"
)

type Handler struct {
	authService auth.AuthService
}

func NewHandler(service auth.AuthService) *Handler {
	return &Handler{
		authService: service,
	}
}

func (h *Handler) Routes(group *huma.Group) {
	user := huma.NewGroup(group, "/auth")

	h.registerLoginAccount(user)

	account := huma.NewGroup(user, "/account")

	h.registerPostAccount(account)
	h.registerPatchAccount(account)
	h.registerDeleteAccount(account)
}
