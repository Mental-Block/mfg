package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/server/env"
	"github.com/server/internal/core/services/auth"
	"github.com/server/internal/core/services/user"
)

type Services struct {
	authService auth.AuthService
	userService user.UserService
}

type App struct {
	services *Services
	router   *chi.Mux
	server   *http.Server
}

const docs = "/docs"

func NewApp(authService auth.AuthService, userService user.UserService, opts ...AppOption) *App {
	cfg := env.Env()
	router := chi.NewMux()

	// should add as options for app config
	adapter := humachi.NewAdapter(router)
	humaConfig := huma.DefaultConfig("MFG", "1.0.0")
	humaConfig.DocsPath = docs
	apiConfig := huma.NewAPI(humaConfig, adapter)

	api := huma.NewGroup(apiConfig, "/api")

	app := &App{
		services: &Services{
			authService: authService,
			userService: userService,
		},
		router: router,
		server: &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf("%s:%s", cfg.Web.Host, cfg.Web.Port),
		},
	}

	for _, applyOption := range opts {
		applyOption(app)
	}

	app.v1(api)

	return app
}

func (a *App) Run() error {
	fmt.Printf("listening at: http://%v%v", a.server.Addr, docs)
	return a.server.ListenAndServe()
}

func (a *App) Stop() error {
	return a.server.Shutdown(context.Background())
}
