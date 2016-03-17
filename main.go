package main

import (
    "fmt"
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
    fmt.Fprintf(writer, "Task Index")
}

func TaskShow(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    fmt.Fprintf(writer, "Task Show:", params["id"])
}
