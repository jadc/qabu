package main

import (
    "context"
    "fmt"
    "os"
    //    "net/http"
    //    "database/sql"
    "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    // Retrieve POSTGRES_URL from the environment
    database_url := os.Getenv("POSTGRES_URL")
    if database_url == "" {
        fmt.Println("Environmental variable POSTGRES_URL is not set")
        os.Exit(1)
    }

    // Connect to PostgreSQL
    db, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
        os.Exit(1)
    }
    defer db.Close()

    // Query the database
    var greeting string
    err = db.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Println(greeting)
}
