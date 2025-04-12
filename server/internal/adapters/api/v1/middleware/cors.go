package middleware

import (
	"github.com/danielgtaylor/huma/v2"
)

func CORS() func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		ctx.AppendHeader("Access-Control-Allow-Credentials", "true")
		ctx.AppendHeader("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.AppendHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.AppendHeader("Access-Control-Allow-Origin", "*")

		next(ctx)
	}
}
