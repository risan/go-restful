package main

import (
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    for _, route := range routes {
        routeHandler := Logger(route.Handler, route.Name)

        router.Methods(route.Method).
               Path(route.Pattern).
               Handler(routeHandler).
               Name(route.Name)
    }

    return router
}
