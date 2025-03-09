package main

import (
	"fmt"
	"os"
    "net/http"
    "github.com/gin-gonic/gin"

    "github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"

    _ "github.com/danielgtaylor/huma/v2/formats/cbor"
)


// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {

        cfg, err := LoadConfig()

        if err != nil {
            fmt.Fprintf(os.Stderr, "config failed: %v\n", err)
        }

		
        router := gin.Default()
		
        api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

        
		// Add the operation handler to the API.
		huma.Get(api, "/greeting/{name}", func(ctx context.Context, input *struct{
			Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
		}) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		})

        // Run the CLI. When passed no commands, it starts the server.
            cli.Run()

        
            // router.GET("/albums", getAlbums)
            // router.GET("/albums/:id", getAlbumByID)
            // router.POST("/albums", postAlbums)

            // url := cfg.Web.Host + ":" + cfg.Web.Port;

            // router.Run(url)


            // Options for the CLI. Pass `--port` or set the `SERVICE_PORT` env var.

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
