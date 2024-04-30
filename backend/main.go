package main

import (
    "context"
    "fmt"
    "os"
    //    "net/http"
    "github.com/jadc/qabu/internal/database"
)

func main() {
    db, err := database.Get()

    _, err = db.Exec(context.Background(), "INSERT INTO greetings (greeting) VALUES ('yo')")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to insert row: %v\n", err)
        os.Exit(1)
    }

    // Query the database
    var greeting string
    err = db.QueryRow(context.Background(), "SELECT * FROM greetings").Scan(&greeting)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Println(greeting)
}
