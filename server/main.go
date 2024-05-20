package main

import (
    "os"
	"net/http"
	"log"
	"github.com/jadc/qabu/internal/api"
)

func logger(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
        h.ServeHTTP(w, r)
    })
}


func main() {
    log.Println("Preparing server")

    // Get the port from the environment, or default to 8080
    var port string = os.Getenv("SERVER_PORT")
    if port == "" {
        log.Println("SERVER_PORT environment variable is not set, defaulting to 8080")
        port = "8080"
    }
    port = ":" + port

    // Create a new router and server
	router := http.NewServeMux()
	server := http.Server{ Addr: port, Handler: logger(router) }

    log.Println("Configuring routes")
    router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, world!"))
    })
	router.HandleFunc("GET /files", api.GetFiles)

	/*
		err = pg.AddFile(ctx, database.File{Title: "test"})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to insert row: %v\n", err)
			os.Exit(1)
		}
	*/

    log.Println("Starting server on port", port)
    log.Fatal(server.ListenAndServe())
}
