package main

import (
    "encoding/json"
    "net/http"
)

type SuccessResponse struct {
    Message string `json:"message"`
    StatusCode int `json:"status_code"`
}

type DataResponse struct {
    Data interface{} `json:"data"`
    StatusCode int `json:"status_code"`
}

type ErrorResponse struct {
    Error string `json:"error"`
    StatusCode int `json:"status_code"`
}

func RespondWithSuccess(writer http.ResponseWriter, message string, statusCode int) {
    response := SuccessResponse{message, statusCode}

    Respond(writer, response, statusCode)
}

func RespondWithData(writer http.ResponseWriter, data interface{}, statusCode int) {
    response := DataResponse{data, statusCode}

    Respond(writer, response, statusCode)
}

func RespondWithError(writer http.ResponseWriter, error string, statusCode int) {
    response := ErrorResponse{error, statusCode}

    Respond(writer, response, statusCode)
}

func Respond(writer http.ResponseWriter, data interface{}, statusCode int) {
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    writer.WriteHeader(statusCode)
    json.NewEncoder(writer).Encode(data)
}
