package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/server/internal/core/ports"
)

type Services struct {
	AuthService  ports.AuthService
	UserService  ports.UserService
	TokenService ports.TokenService
}

type API struct {
	port     int
	host     string
	name     string
	version  string
	docsPath string
	router   http.Handler
	server   http.Server /* non configurable dependancy */
	services Services    /* non configurable dependancy */
}

func NewAPI(services Services, opts ...APIOption) *API {
	api := &API{
		version:  "1.0.0",
		name:     "MyAPI",
		docsPath: "/docs",
		router:   http.NewServeMux(),
		host:     "localhost",
		port:     8084,
	}

	for _, applyOption := range opts {
		applyOption(api)
	}

	api.services = services

	api.server = http.Server{
		Addr:    fmt.Sprintf("%v:%v", api.host, api.port),
		Handler: api.router,
	}

	if _, ok := api.router.(*chi.Mux); ok && api.version == "1.0.0" {
		api.v1()
	} else {
		slog.Error("not a valid api version or supported mux. Current versions include: \n 1.0.0: chi.Mux")
	}

	return api
}

func (a *API) Run() error {
	slog.Info(fmt.Sprintf("listening at: http://%v:%v%v", a.host, a.port, a.docsPath))
	return a.server.ListenAndServe()
}

func (a *API) Stop() error {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)

	defer cancel()

	return a.server.Shutdown(ctx)
}
