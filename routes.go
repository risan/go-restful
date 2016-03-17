package main

import (
    "net/http"
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
