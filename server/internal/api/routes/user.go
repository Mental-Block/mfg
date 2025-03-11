package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func user() chi.Router {
    r := chi.NewRouter()
	
	
    r.Route("/hello", func(r chi.Router){
        r.Get("/", hello)
    })

    return r
}


func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}