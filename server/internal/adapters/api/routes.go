package api

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/server/internal/adapters/api/v1/handler/authentication"
	"github.com/server/internal/adapters/api/v1/handler/permission"
	"github.com/server/internal/adapters/api/v1/handler/resource"
	"github.com/server/internal/adapters/api/v1/handler/role"
	"github.com/server/internal/adapters/api/v1/handler/user"
	"github.com/server/internal/core/domain"
)

func (a *API) v1() {
	humaConfig := huma.DefaultConfig(a.name, a.version)
	humaConfig.DocsPath = a.docsPath
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
		domain.RefreshTokenName: {
			Description:  "long lived token to request auth tokens",
			In: 		  "cookie",	
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
		domain.AuthTokenName: {
			Description:  "short lived auth token with user roles",
			In: 		  "cookie",
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	

	humaBaseAPI := humachi.New(a.router.(*chi.Mux), humaConfig)

	api := huma.NewGroup(humaBaseAPI, "/api")

	v1 := huma.NewGroup(api, "/v1")
	
	user.NewUserHandler(a.services.UserService).Routes(v1)
	authentication.NewAuthHandler(a.services.AuthService, a.services.UserService).Routes(v1)
	role.NewRoleHandler(a.services.RoleService).Routes(v1)
	permission.NewPermissionHandler(a.services.PermissionService).Routes(v1)
	resource.NewResourceHandler(a.services.ResourceService).Routes(v1)
}
