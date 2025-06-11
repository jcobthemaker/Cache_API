package main

import (
	"cache_API/db"
	"context"
	"fmt"
	"net/http"
)

var ctx = context.Background()

func main() {
    InitRedis()
    database, err := db.InitDB()

    if err != nil {
        fmt.Println(err)
    }

        defer func() {
            database.Close()
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
            if err, ok := r.(error); ok {
                fmt.Printf("Error: %v\n", err)
            }
        }
    }()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        HandleRequest(w, r, ctx, database)
    })

    http.ListenAndServe(":8080", nil)
}