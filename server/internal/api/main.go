package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/server/env"
	"github.com/server/internal/api/routes"
	"github.com/server/internal/postgres"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
)

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *struct {}) {
		cfg := env.Env(); 

		conn, err := postgres.NewRepo(context.Background(), cfg.DB.URL);

		var router = chi.NewMux();

		routes.Routes(router)

		server := &http.Server {
			Addr: fmt.Sprintf("%s:%s", cfg.Web.Host, cfg.Web.Port),
			Handler: router,
		}
		
		hooks.OnStart(func() {
			fmt.Printf("debug:%v host:%v port%v\n",
			"true",  cfg.Web.Host, cfg.Web.Port)

			if err != nil {
				fmt.Printf("database failed: %s\n", err)
			}

			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("web server failed: %s\n", err)
			}
		})

		hooks.OnStop(func () {	
			conn.Close()
			server.Shutdown(context.Background())
		})
	})

	cli.Run()
}