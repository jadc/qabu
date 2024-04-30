package database

import (
    "context"
    "fmt"
    "os"
    "github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func init() {
    // Retrieve POSTGRES_URL from the environment
    database_url := os.Getenv("POSTGRES_URL")
    if database_url == "" {
        fmt.Fprintf(os.Stderr, "Environmental variable POSTGRES_URL is not set")
        os.Exit(1)
    }

    // Connect to PostgreSQL
    var err error
    db, err = pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to connect to PostgreSQL: %v\n", err)
        os.Exit(1)
    }

    createTables(db)
}

func Get() (*pgxpool.Pool, error) {
    return db, nil
}
