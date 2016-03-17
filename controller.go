package main

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Homepage")
}

func TaskIndex(writer http.ResponseWriter, request *http.Request) {
    tasks := Tasks{
        Task{Name: "Learn Golang"},
        Task{Name: "Build RESTful API with Golang"},
    }

    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(tasks)
}

func TaskShow(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    fmt.Fprintf(writer, "Task Show:", params["id"])
}
