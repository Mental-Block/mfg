package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type APIOption func(a *API)

func WithRouter(router http.Handler) APIOption {
	return func(a *API) {
		a.router = router
	}
}

func WithDocs(path string) APIOption {
	return func(a *API) {
		a.docsPath = path
	}
}

func WithVersion(version string) APIOption {
	return func(a *API) {
		a.version = version
	}
}

func WithName(name string) APIOption {
	return func(a *API) {
		a.name = name
	}
}

func WithPort(port string) APIOption {
	return func(a *API) {
		port, err := strconv.Atoi(port)

		if err != nil {
			slog.Error(fmt.Sprintf("%v is invalid port number. falling back to defualt %v", port, a.port))
			return
		}

		a.port = port
	}
}

func WithHost(host string) APIOption {
	return func(a *API) {
		a.host = host
	}
}
