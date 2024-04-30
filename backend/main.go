package main

import (
    "fmt"
    "os"
    //    "net/http"
    //    "database/sql"
)

func main() {
    database_url := os.Getenv("DATABASE_URL")
    if database_url == "" {
        fmt.Println("Environmental variable DATABASE_URL is not set")
        os.Exit(1)
    }

    fmt.Println(database_url)
}
