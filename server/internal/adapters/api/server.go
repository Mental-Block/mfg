package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/server/internal/core/ports"
)

type Services struct {
	AuthService  		ports.AuthService
	UserService  		ports.UserService
	RoleService     	ports.RoleService
	PermissionService 	ports.PermissionService
	ResourceService 	ports.ResourceService
}

type API struct {
	port     	int
	host     	string
	name     	string
	version  	string
	docsPath 	string
	router   	http.Handler
	middleware 	[]func(next http.Handler) http.Handler
	services 	Services
	Server   	http.Server     
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

	var handler http.Handler = api.router
	for _, applyMiddleware := range api.middleware {
		handler = applyMiddleware(handler)
	}

	api.services = services

	api.Server = http.Server{
		Addr:    		  fmt.Sprintf("%v:%v", api.host, api.port),
		Handler: 		  handler,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}

	if _, ok := api.router.(*chi.Mux); ok && api.version == "1.0.0" {
		api.v1()
	} else {
		slog.Error("not a valid api version or supported mux. Current versions include: \n 1.0.0: chi.Mux")
	}

	return api
}
