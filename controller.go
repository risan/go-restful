package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
)

func Index(writer http.ResponseWriter, request *http.Request) {
    RespondWithSuccess(writer, "RESTful API with Golang", http.StatusOK)
}

func TaskIndex(writer http.ResponseWriter, request *http.Request) {
    tasks, error := GetAllTasks()

    if error != nil {
        RespondWithError(writer, "An error has occured.", http.StatusInternalServerError)
        return
    }

    RespondWithData(writer, tasks, http.StatusOK)
}

func TaskShow(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)

    var task Task

    error := FindTaskById(params["id"], &task)

    switch {
        case error == mgo.ErrNotFound:
            RespondWithError(writer, "Resource not found.", http.StatusNotFound)
            return
        case error != nil:
            RespondWithError(writer, "An error has occured.", http.StatusInternalServerError)
            return
    }

    RespondWithData(writer, task, http.StatusOK)
}
