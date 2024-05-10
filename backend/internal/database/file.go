package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// Simplified representation of a file for listing
// The rest of the metadata is stored in other tables
type File struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func init() {
	// Connect to database
	pg, err := Connect(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Create table
	query := `
        CREATE TABLE IF NOT EXISTS files (
            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            title TEXT NOT NULL
        )
    `
	_, err = pg.db.Exec(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create table: %v\n", err)
		os.Exit(1)
	}
}

func (pg *postgres) ListFiles(ctx context.Context) ([]File, error) {
	query := `SELECT * FROM files`

	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to query: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[File])
}

func (pg *postgres) AddFile(ctx context.Context, file File) error {
	query := `INSERT INTO files (title) VALUES ($1)`
	_, err := pg.db.Exec(ctx, query, file.Title)
	return err
}
