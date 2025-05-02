package main

import (
	"context"
	"log"
	"net/http"

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

		logger.Set(cfg.ENV)

		authJwt := jwt.New(cfg.API.AuthSecret)
		emailJwt := jwt.New(cfg.API.EmailSecret)
		argon := argon.New(cfg.API.PasswordSalt)

		postgres, postgresErr := postgres.NewStore(context.Background(), cfg.DB.URL)
		redis, redisErr := redis.NewStore(context.Background(), cfg.DB_CACHE.URL)
		smtp := smtp.NewSMTP(cfg.SMTP.HostEmail, cfg.SMTP.Password, cfg.SMTP.Host, cfg.SMTP.Port)

		tokenStore := store.NewTokenStore(redis)
		authStore := store.NewAuthStore(postgres, redis)
		userStore := store.NewUserStore(postgres)
		userProfileStore := store.NewUserProfileStore(postgres)
		
		authService := services.NewAuthService(
			userStore,
			authStore,
			tokenStore,
			argon,
			smtp,
			authJwt,
			emailJwt,
		)
		
		userService := services.NewUserService(
			userStore,
			userProfileStore,
		)

		server := api.NewAPI(
			api.Services{
				AuthService:  authService,
				UserService:  userService,
			},
			api.WithDocs("/docs"),
			api.WithHost(cfg.API.Host),
			api.WithPort(cfg.API.Port),
			api.WithVersion("1.0.0"),
			api.WithName("MFG"),
			api.WithRouter(chi.NewMux()),
		)

		hooks.OnStart(func() {
			if redisErr != nil {
				log.Print("failed to init cache database: %w\n", redisErr)
			}

			if postgresErr != nil {
				log.Print("failed to init database: %w\n", postgresErr)
			}

			if err := server.Run(); err != nil && err != http.ErrServerClosed {
				log.Print("failed to init web server: %w\n", err)
			}
		})

		hooks.OnStop(func() {
			if err := server.Stop(); err != nil {
				log.Print("web could not shut down gracefully: %w\n", err)
			}

			if err := redis.Close(); err != nil {
				log.Print("cache db couldn't close connection gracefully: %w\n", err)
			}

			postgres.Close()
		})
	})

	cli.Run()
}
