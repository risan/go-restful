package main

import "gopkg.in/mgo.v2/bson"

type Task struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Name string `json:"name"`
    IsCompleted bool `json:"is_completed"`
}

type Tasks []Task
