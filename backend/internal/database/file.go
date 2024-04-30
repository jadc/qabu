package database

import (
    "context"
    "fmt"
    "os"
    "github.com/jackc/pgx/v5"
)

type File struct {
    FilePath string
    MD5 string
}

func init() {
    // Connect to database
    pg, err := Connect(context.Background())

    // Create table
    query := `
        CREATE TABLE IF NOT EXISTS files (
            file_path text, 
            md5 text
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
    query := `INSERT INTO files (file_path, md5) VALUES ($1, $2)`
    _, err := pg.db.Exec(ctx, query, file.FilePath, file.MD5)
    return err
}
