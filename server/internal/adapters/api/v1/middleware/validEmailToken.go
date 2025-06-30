package middleware

// import (
// 	"github.com/danielgtaylor/huma/v2"
// 	"github.com/server/internal"
// )

// func ValidEmailTokenMiddleWare(ctx huma.Context, next func(huma.Context)) {

// 	// if ctx.Query("error") == "true" {
// 	// 	huma.WriteErr(api, ctx, http.StatusInternalServerError,
// 	// 		"Some friendly message", fmt.Errorf("error detail"),
// 	// 	)
// 	// 	return
// 	// }

// 		// Set a custom header on the response.
// 	ctx.SetHeader("My-Custom-Header", "Hello, world!")

// 	claims, err := s.emailTokenService.Parse(token)

// 	if (err != nil) {
// 		huma.Error401Unauthorized("token not valid")

// 		internal.NewErrorf(internal.ErrorCodeNotAuthorized, )
// 	}

// 	email, ok := claims["email"].(string)

// 	if !ok {
// 		return internal.NewErrorf(internal.ErrorCodeUnknown, "could not convert to string")
// 	}

// 	next(ctx)
// }

// // NewAuthMiddleware creates a middleware that will authorize requests based on
// // the required scopes for the operation.
// func NewAuthMiddleware(api huma.API, jwksURL string) func(ctx huma.Context, next func(huma.Context)) {
// 	keySet := NewJWKSet(jwksURL)

// 	return func(ctx huma.Context, next func(huma.Context)) {
// 		var anyOfNeededScopes []string
// 		isAuthorizationRequired := false
// 		for _, opScheme := range ctx.Operation().Security {
// 			var ok bool
// 			if anyOfNeededScopes, ok = opScheme["myAuth"]; ok {
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

// 		// Parse and validate the JWT.
// 		parsed, err := jwt.ParseString(token,
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