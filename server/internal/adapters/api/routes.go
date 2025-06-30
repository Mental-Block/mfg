package api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/server/internal/adapters/api/v1/handler/authentication"
	"github.com/server/internal/adapters/api/v1/handler/user"
)

func (a *Config) v1() {
	router := a.router.(*chi.Mux)
	humaConfig := huma.DefaultConfig(a.Name, a.Version)
	humaConfig.DocsPath = a.DocsPath
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		// "google": {
		// 	Type: "oauth2",
		// 	Flows: &huma.OAuthFlows{
		// 		AuthorizationCode: &huma.OAuthFlow{
		// 			AuthorizationURL: "https://example.com/oauth/authorize",
		// 			TokenURL:         "https://example.com/oauth/token",
		// 			Scopes: map[string]string{
		// 				"scope1": "Scope 1 description...",
		// 				"scope2": "Scope 2 description...",
		// 			},
		// 		},
		// 	},
		// },
		// auth_domain.RefreshTokenName: {
		// 	Description:  "long lived, token to request auth tokens",
		// 	In: 		  "cookie",	
		// 	Type:         "http",
		// 	Scheme:       "bearer",
		// 	BearerFormat: "JWT",
		// },
		// auth_domain.AuthTokenName: {
		// 	Description:  "short lived, token to request auth permissions",
		// 	In: 		  "cookie",
		// 	Type:         "http",
		// 	Scheme:       "bearer",
		// 	BearerFormat: "JWT",
		// },
	}

	humaBaseAPI := humachi.New(router, humaConfig)

	api := huma.NewGroup(humaBaseAPI, "/api")

	switch a.DocUI {
		case "scalar":
			router.Get(a.DocsPath, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "text/html")
					w.Write([]byte(`<!doctype html>
				<html>
				<head>
					<title>API Reference</title>
					<meta charset="utf-8" />
					<meta
					name="viewport"
					content="width=device-width, initial-scale=1" />
				</head>
				<body>
					<script
					id="api-reference"
					data-url="/openapi.json"></script>
					<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
				</body>
				</html>`))
			})
		case "swagger":
			router.Get(a.DocsPath, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(`<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charset="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<meta name="description" content="SwaggerUI" />
				<title>SwaggerUI</title>
				<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
				</head>
				<body>
				<div id="swagger-ui"></div>
				<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
				<script>
				window.onload = () => {
					window.ui = SwaggerUIBundle({
					url: '/openapi.json',
					dom_id: '#swagger-ui',
					});
				};
				</script>
				</body>
				</html>`))
			})
		default:
	}

	v1 := huma.NewGroup(api, "/v1")
	
	user.NewUserHandler(a.services.UserService).Routes(v1)
	authentication.NewAuthHandler(a.Host, a.environment, a.services.AuthService).Routes(v1)
}
