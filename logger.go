package main

import (
    "log"
    "net/http"
    "time"
)

func Logger(routeHandler http.Handler, routeName string) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        start := time.Now()

        routeHandler.ServeHTTP(writer, request)

        log.Printf(
            "[%s]\t%s\t%s\t[%s]",
            request.Method,
            request.RequestURI,
            routeName,
            time.Since(start),
        )
    })
}
