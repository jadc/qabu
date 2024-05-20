package main

import (
    "os"
	"net/http"
	"log"
	"github.com/jadc/qabu/internal/api"
)

func newRouter() *http.ServeMux {
    router := http.NewServeMux()
    templates := api.GetTemplates()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)

        var data struct {
            Test []int
        }
        data.Test = []int{1, 2, 3, 4}
        templates.Render(w, "index.html", data)
    })

    router.HandleFunc("/files", api.GetFiles)

    return router
}

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
	server := http.Server{ Addr: port, Handler: logger(newRouter()) }

    log.Println("Starting server on port", port)
    log.Fatal(server.ListenAndServe())
}
