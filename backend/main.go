package main

import (
    "context"
    "fmt"
    "os"
    //    "net/http"
    "github.com/jadc/qabu/internal/database"
)

func main() {
    ctx := context.Background()
    pg, err := database.Connect(ctx)

    err = pg.AddFile(ctx, database.File{FilePath: "file1.txt", MD5: "1234567890"})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to insert row: %v\n", err)
        os.Exit(1)
    }

    // Query the database
    files, err := pg.ListFiles(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }

    for _, file := range files {
        fmt.Fprintf(os.Stdout, "file_path: %s, md5: %s\n", file.FilePath, file.MD5)  
    }
}
