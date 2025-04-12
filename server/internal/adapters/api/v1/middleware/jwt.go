package middleware

// import (
// 	"net/http"
// 	"slices"
// 	"strings"

// 	"github.com/danielgtaylor/huma/v2"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/server/internal/core/ports"
// )

// func NewAuthMiddleware(api huma.API, service ports.TokenService) func(ctx huma.Context, next func(huma.Context)) {
// 	//keySet := NewJWKSet(jwksURL)

// 	return func(ctx huma.Context, next func(huma.Context)) {
// 		var anyOfNeededScopes []string
// 		isAuthorizationRequired := false
// 		for _, opScheme := range ctx.Operation().Security {
// 			var ok bool
// 			if anyOfNeededScopes, ok = opScheme["delete-account"]; ok {
// 				isAuthorizationRequired = true
// 				break
// 			}
// 		}

// 		if !isAuthorizationRequired {
// 			next(ctx)
// 			return
// 		}

// 		token := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
// 		if len(token) == 0 {
// 			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Unauthorized")
// 			return
// 		}

// 		jwt.Parse()

// 		// Parse and validate the JWT.
// 		parsed, err := jwt.Parse(token,
// 			jwt.WithKeySet(keySet),
// 			jwt.WithValidate(true),
// 			jwt.WithIssuer("my-issuer"),
// 			jwt.WithAudience("my-audience"),
// 		)
// 		if err != nil {
// 			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Unauthorized")
// 			return
// 		}

// 		// Ensure the claims required for this operation are present.
// 		scopes, _ := parsed.Get("scopes")
// 		if scopes, ok := scopes.([]string); ok {
// 			for _, scope := range scopes {
// 				if slices.Contains(anyOfNeededScopes, scope) {
// 					next(ctx)
// 					return
// 				}
// 			}
// 		}

// 		huma.WriteErr(api, ctx, http.StatusForbidden, "Forbidden")
// 	}
// }
