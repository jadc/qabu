package main

import (
	"net/http"

	"github.com/jadc/qabu/internal/api"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{Addr: ":8080", Handler: router}

	router.HandleFunc("GET /files", api.GetFiles)

	/*
		err = pg.AddFile(ctx, database.File{Title: "test"})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to insert row: %v\n", err)
			os.Exit(1)
		}
	*/

	server.ListenAndServe()
}
