package database

import (
	"context"
	"log"
	"os/exec"
	"encoding/json"
)

type File struct {
    UUID      string
    FileName  string
    Size      int64
    Type      string
    Tags      []string
    Created   string
    Leaked    string
    Exif      string
    Original  bool
}

func init() {
	// Connect to database
	pg, err := Connect(context.Background())
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Create table
	query := `
        CREATE TABLE IF NOT EXISTS files (
            uuid UUID PRIMARY KEY,
            file_name TEXT NOT NULL,
            size BIGINT,
            type TEXT,
            tags TEXT[],
            created TEXT,
            leaked TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            exif JSONB NOT NULL,
            original BOOLEAN DEFAULT FALSE,
        )
    `
	_, err = pg.db.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}

/*
func (pg *postgres) ListFiles(ctx context.Context) ([]File, error) {
	query := `SELECT * FROM files`

	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to query: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[File])
}
*/

// Given a file struct, save it to the database
func (pg *postgres) SaveFile(ctx context.Context, file File) error {
	query := `INSERT INTO files (uuid, file_name, size, type, tags, created, leaked, exif, original) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := pg.db.Exec(ctx, query, file.UUID, file.FileName, file.Size, file.Type, file.Tags, file.Created, file.Leaked, file.Exif, file.Original)
	return err
}

// Given a UUID, delete the file from the database
func (pg *postgres) DeleteFile(ctx context.Context, uuid string) error {
    query := `DELETE FROM files WHERE uuid = $1`
    _, err := pg.db.Exec(ctx, query, uuid)
    return err
}

// Given a local path to a file, create a new file struct and pre-fill some fields based on the file's exif data, then return the struct
func (pg *postgres) CreateFile(ctx context.Context, path string) (*File, error) {
    new_file := File{
        UUID: "",
        FileName: "",
        Size: 0,
        Type: "",
        Tags: []string{},
        Created: "",
        Leaked: "",
        Exif: "",
        Original: false,
    }

    // Get checksum
    sha_cmd := exec.Command("sha256sum", path)
    sha_stdout, err := sha_cmd.Output()
    if err != nil {
        log.Println("Failed to get checksum: ", err)
        return nil, err
    }
    new_file.UUID = string(sha_stdout[:64])

    // Get exif data
    cmd := exec.Command("exiftool", "-n", "-json", path)
    stdout, err := cmd.Output()
    if err != nil {
        log.Println("Failed to get exif data: ", err)
        return nil, err
    }

    // Parse exif data
    exif := make(map[string]interface{})
    err = json.Unmarshal(stdout, &exif)
    if err != nil {
        log.Println("Failed to parse exif data: ", err)
        return nil, err
    }

    // Fill in file struct with exif data when possible
    new_file.FileName = exif["FileName"].(string)
    new_file.Size = exif["FileSize"].(int64)
    new_file.Type = exif["FileType"].(string)

    date_labels := []string{"FileModifyDate", "FileAccessDate", "FileInodeChangeDate", "FilePermissionsModifyDate", "FileCreationDate", "Year", "DateTimeOriginal"}
    dates := []string{}

    // Get dates
    for _, date := range date_labels {
        if exif[date] != nil {
            dates = append(dates, exif[date].(string))
        }
    }

    for _, date := range dates {
        // Find the earliest date
        if new_file.Created > date {
            new_file.Created = date
        }

        // Find the latest date
        if new_file.Leaked < date {
            new_file.Leaked = date
        }
    }

    new_file.Exif = string(stdout)

    return &new_file, err
}
