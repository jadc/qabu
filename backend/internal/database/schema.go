package database

import (
    "context"
    "fmt"
    "os"
    "github.com/jackc/pgx/v5/pgxpool"
)

func createTables(db *pgxpool.Pool) {
    _, err := db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS greetings (greeting text)")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create table: %v\n", err)
        os.Exit(1)
    }
}

