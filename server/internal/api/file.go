package api

// TODO: replace with posts api, not files

/*
import (
	"context"
	"encoding/json"
	"net/http"
	"log"
	"github.com/jadc/qabu/internal/database"
)

func GetFiles(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	pg, err := database.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Query the database
	files, err := pg.ListFiles(ctx)
	if err != nil {
		log.Fatal("Failed to query database: ", err)
	}

	res, err := json.Marshal(files)
	if err != nil {
		log.Fatal("Failed to marshal response: ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// TODO
	w.WriteHeader(http.StatusOK)
	//w.Write(res)
}
*/
