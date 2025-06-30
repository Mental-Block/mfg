package api

import (
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/server/config/env"
	"github.com/server/internal/core/auth"
	"github.com/server/internal/core/user"
)

type Services struct {
	AuthService  		auth.AuthService
	UserService  		user.UserService
}

type Config struct {
	Port     	int		`yaml:"port" default:"8080"`
	Host     	string	`yaml:"host" default:"localhost"`
	Name     	string	`yaml:"name" default:"mfg"`
	Version  	string	`yaml:"version" default:"1.0.0"`
	DocsPath 	string	`yaml:"doc_path" default:"/docs"`
	DocUI		string	`yaml:"doc_ui" default:"stoplight"`
	CorsConfig  CorsConfig `yaml:"cors"`
	router   	http.Handler `yaml:"-"`
	environment env.ENVIROMENT `yaml:"-"`
	middleware 	[]func(next http.Handler) http.Handler `yaml:"-"`
	services 	Services `yaml:"-"`
	Server   	http.Server `yaml:"-"`     
}

func Serve(services Services, opts ...APIOption) *Config {
	api := &Config{
		Version:  "1.0.0",
		Name:     "MyAPI",
		DocUI:    "stoplight",
		DocsPath: "/docs",
		router:   chi.NewMux(),
		Host:     "localhost",
		Port:     8084,
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
		Addr:    		  net.JoinHostPort(api.Host, strconv.Itoa(api.Port)),
		Handler: 		  handler,
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       2 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	if _, ok := api.router.(*chi.Mux); ok && api.Version == "1.0.0" {
		api.v1()
	} else {
		slog.Error("not a valid api version or supported mux. Current versions include: \n 1.0.0: chi.Mux")
	}

	return api
}
