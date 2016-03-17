package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Route struct {
    Method string
    Pattern string
    Handler http.HandlerFunc
    Name string
}

type Routes []Route

var routes = Routes{
    Route{"GET", "/", Index, "index"},
    Route{"GET", "/tasks", TaskIndex, "task.index"},
    Route{"GET", "/tasks/{id}", TaskShow, "task.show"},
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    for _, route := range routes {
        router.Methods(route.Method).
               Path(route.Pattern).
               Handler(route.Handler).
               Name(route.Name)
    }

    return router
}
