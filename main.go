package main

import (
    "log"
    "net/http"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), router))
}
