package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", Index)
    router.HandleFunc("/tasks", TaskIndex)
    router.HandleFunc("/tasks/{id}", TaskShow)

    log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Homepage")
}

func TaskIndex(writer http.ResponseWriter, request *http.Request) {
    tasks := Tasks{
        Task{Name: "Learn Golang"},
        Task{Name: "Build RESTful API with Golang"},
    }

    json.NewEncoder(writer).Encode(tasks)
}

func TaskShow(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    fmt.Fprintf(writer, "Task Show:", params["id"])
}
