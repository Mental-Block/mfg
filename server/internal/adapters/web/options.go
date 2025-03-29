package web

import (
	"net/http"
)

type AppOption func(a *App)

func WithServer(server *http.Server) AppOption {
	return func(a *App) {
		a.server = server
	}
}
