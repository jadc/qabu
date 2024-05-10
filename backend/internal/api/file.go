package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jadc/qabu/internal/database"
)

func GetFiles(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	pg, err := database.Connect(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Query the database
	files, err := pg.ListFiles(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	res, err := json.Marshal(files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Marshaling failed: %v\n", err)
		os.Exit(1)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
