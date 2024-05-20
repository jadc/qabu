package main

import (
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

const PORT string = ":8000"

func main() {
    log.Println("Preparing server")
	router := http.NewServeMux()
	server := http.Server{ Addr: PORT, Handler: logger(router) }

    log.Println("Configuring routes")
    router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, World!"))
    })
	router.HandleFunc("GET /files", api.GetFiles)

	/*
		err = pg.AddFile(ctx, database.File{Title: "test"})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to insert row: %v\n", err)
			os.Exit(1)
		}
	*/

    log.Println("Starting server on port", PORT)
    log.Fatal(server.ListenAndServe())
}
