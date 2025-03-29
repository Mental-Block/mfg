package main

import (
	"context"
	"log"
	"net/http"

	"github.com/server/env"
	"github.com/server/internal/adapters/store"
	"github.com/server/internal/adapters/web"
	"github.com/server/internal/adapters/web/cli"
	"github.com/server/internal/core/services/auth"
	"github.com/server/internal/core/services/user"
)

func main() {
	cfg := env.Env()

	conn, err := store.NewStore(context.Background(), cfg.DB.URL)

	authService := auth.NewService(conn)
	userService := user.NewService(conn)

	srv := web.NewApp(authService, userService)

	cli := cli.New(func() {

		if err != nil {
			log.Fatal("failed to init database failed: %w\n", err)
		}

		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to init web server failed: %w\n", err)
		}

	},
		func() {
			conn.Close()
			if err := srv.Stop(); err != nil {
				log.Fatal("web could not shut down gracefully: %w\n", err)
			}
		})

	cli.Run()
}
