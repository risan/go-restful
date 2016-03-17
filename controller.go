package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

func Index(writer http.ResponseWriter, request *http.Request) {
    RespondWithSuccess(writer, "RESTful API with Golang", http.StatusOK)
}

func TaskIndex(writer http.ResponseWriter, request *http.Request) {
    tasks := Tasks{
        Task{Name: "Learn Golang language"},
        Task{Name: "Build RESTful API"},
    }

    RespondWithData(writer, tasks, http.StatusOK)
}

func TaskShow(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)

    task := Task{Name: "Task id: " + params["id"]}

    RespondWithData(writer, task, http.StatusOK)
}
