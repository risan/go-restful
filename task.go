package main

import "time"

type Task struct {
    Name string `json:"name"`
    IsCompleted bool `json:"is_completed"`
    Due time.Time `json:"time"`
}

type Tasks []Task
