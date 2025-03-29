package web

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/web/auth"
	"github.com/server/internal/adapters/web/user"
)

func (a *App) v1(api huma.API) {
	v1 := huma.NewGroup(api, "/v1")

	user.NewHandler(a.services.userService).Routes(v1)
	auth.NewHandler(a.services.authService).Routes(v1)
}
