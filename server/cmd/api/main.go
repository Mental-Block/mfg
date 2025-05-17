package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"runtime"
	"time"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"

	"github.com/server/internal/adapters/api"
	"github.com/server/internal/adapters/auth/argon"
	"github.com/server/internal/adapters/auth/jwt"
	"github.com/server/internal/adapters/env"
	"github.com/server/internal/adapters/logger"
	"github.com/server/internal/adapters/smtp"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/adapters/store/redis"
	"github.com/server/internal/adapters/store/store"
	"github.com/server/internal/core/services"
)

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *struct{}) {

		cfg := env.Env()
		
		// Primary Adapters
		// token
		authJwt := jwt.New(cfg.API.AuthSecret)
		emailJwt := jwt.New(cfg.API.EmailSecret)
		refreshJwt := jwt.New(cfg.API.RefreshSecret)

		// password
		argon := argon.New(64 * 1024, 1,uint8(runtime.NumCPU()), 16, 32)
		
		// RDBMS
		postgres, postgresErr := postgres.NewStore(context.Background(), cfg.DB.URL)
		
		// In-Memory NoSQL (cache)
		redis, redisErr := redis.NewStore(context.Background(), cfg.DB_CACHE.URL)
		
		// data stores/repos
		tokenStore := store.NewTokenStore(redis)

		authStore := store.NewAuthStore(postgres, redis)
		authUserStore := store.NewAuthUserStore(postgres)

		userStore := store.NewUserStore(postgres)
		//userAuthStore := store.NewUserAuthStore(postgres)

		roleStore := store.NewRoleStore(postgres)
		permissionStore := store.NewPermissionStore(postgres)
		resourceStore := store.NewResourceStore(postgres)

		
		// SMTP
		smtp := smtp.NewSMTP(cfg.SMTP.HostEmail, cfg.SMTP.Password, cfg.SMTP.Host, cfg.SMTP.Port)

		// Secondary Adapters
		authService := services.NewAuthService(
			authUserStore,
			userStore,
			authStore,
			tokenStore,
			argon,
			smtp,
			refreshJwt,
			authJwt,
			emailJwt,
		)
		userService := services.NewUserService(userStore)
		roleService := services.NewRoleService(roleStore)
		permissionService := services.NewPermissionService(permissionStore)
		resourceService := services.NewResourceService(resourceStore)
		
		//logging middleware
		logger.Set(cfg.ENV)

		logging := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				slog.Info(
					r.Method,
					slog.Time("time", time.Now()),
					slog.String("url", r.URL.String()),
				)

				h.ServeHTTP(w, r)
			})
		}
 
		//cors middleware
		cors := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				
				if r.Method == http.MethodOptions && r.Header.Get("Origin") != "" && r.Header.Get("Access-Control-Request-Method") != "" {
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
					w.Header().Set("Access-Control-Max-Age", "600")

					w.WriteHeader(204)
					return
				}

				w.Header().Add("Vary", "Origin")
				h.ServeHTTP(w, r)
			})
		}

		// REST API
		server := api.NewAPI(
			api.Services{
				AuthService:  authService,
				UserService:  userService,
				RoleService:  roleService,
				PermissionService: permissionService,
				ResourceService: resourceService,
			},
			api.WithMiddlewares([]func(next http.Handler) http.Handler{logging, cors}),
			api.WithDocs("/docs"),
			api.WithHost(cfg.API.Host),
			api.WithPort(cfg.API.Port),
			api.WithVersion("1.0.0"),
			api.WithName("MFG"),
			api.WithRouter(chi.NewMux()),
		)	
		
		// runs a go routine and SIGTERM process under the hood... so when services start and stop they shutdown gracefully
		hooks.OnStart(func() {
			log.Print("listening at: http://", server.Server.Addr)

			if redisErr != nil {
				log.Print("failed to init cache database: %w\n", redisErr)
			}

			if postgresErr != nil {
				log.Print("failed to init database: %w\n", postgresErr)
			}

			if err := server.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Print("failed to init web server: %w\n", err)
			}
		})

		hooks.OnStop(func() {
			log.Print("shutdown signal recived")
				
			postgres.Close()

			if err := redis.Close(); err != nil {
				log.Print("cache db couldn't close connection gracefully: %w\n", err)
			}

			ctx, cancel := context.WithTimeout(
				context.Background(), time.Duration(3)*time.Second,
			)

			defer cancel()

			if err := server.Server.Shutdown(ctx); err != nil {
				log.Print("web could not shut down gracefully: %w\n", err)
			}
		})
	})

	cli.Run()
}
