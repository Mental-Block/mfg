package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
		Bob string 
	}
}

type GreatingParam struct{
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

type Name struct {
    Last  string `db:"last_name"`
}

func v1() chi.Router {
    r := chi.NewRouter()
    r.Mount("/user", user())
    return r
}

func versions() chi.Router {
    route := chi.NewRouter()
    route.Mount("/v1", v1())
    return route    
}

func Routes(router *chi.Mux) {	
	humaConfig := huma.DefaultConfig("MFG", "1.0.0")
	humaConfig.DocsPath = "/api/v1/"

	router.Mount("/api", versions())
	humachi.New(router, humaConfig)	
}