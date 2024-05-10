package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var pg *postgres

var once sync.Once
var connectionError error

// Connects to the Postgres database and returns a pointer to the postgres struct singleton
func Connect(ctx context.Context) (*postgres, error) {
	// Runs only once
	once.Do(func() {
		// Retrieve POSTGRES_URL from the environment
		database_url := os.Getenv("POSTGRES_URL")
		if database_url == "" {
			connectionError = fmt.Errorf("Environmental variable POSTGRES_URL is not set")
			return
		}

		// Create a connection pool
		db, connectionError := pgxpool.New(ctx, database_url)
		if connectionError != nil {
			return
		}

		// Create new postgres instance with the connection pool
		pg = &postgres{db}
		connectionError = pg.db.Ping(ctx)
	})

	// Returns current or new postgres instance
	return pg, connectionError
}

// Close the connection to the Postgres database
func (pg *postgres) Close() {
	pg.db.Close()
}
