package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/server/config/env"
)

type APIOption func(a *Config)

func WithRouter(router http.Handler) APIOption {
	return func(a *Config) {
		a.router = router
	}
}

func WithDocPath(path string) APIOption {
	return func(a *Config) {
		a.DocsPath = path
	}
}

func WithDocUI(ui string) APIOption {
	return func(a *Config) {
		a.DocUI = ui
	}
}

func WithVersion(version string) APIOption {
	return func(a *Config) {
		a.Version = version
	}
}

func WithName(name string) APIOption {
	return func(a *Config) {
		a.Name = name
	}
}

func WithPort(port string) APIOption {
	return func(a *Config) {
		port, err := strconv.Atoi(port)

		if err != nil {
			slog.Error(fmt.Sprintf("%v is invalid port number. falling back to defualt %v", port, a.Port))
			return
		}

		a.Port = port
	}
}

func WithHost(host string) APIOption {
	return func(a *Config) {
		a.Host = host
	}
}

func WithMiddlewares(middleware []func(next http.Handler) http.Handler) APIOption {
	return func(a *Config) {
		a.middleware = middleware
	}
}

func WithEnivorment(env env.ENVIROMENT) APIOption {
	return func(a *Config) {
		a.environment = env
	}
}