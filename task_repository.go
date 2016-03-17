package main

import (
    "errors"
    "os"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func GetAllTasks() (Tasks, error)  {
    connection, db, error := DbConnection()

    if error != nil {
        return nil, error
    }

    defer connection.Close()

    var tasks Tasks

    error = TaskCollection(db).Find(nil).All(&tasks)

    if error != nil {
        return nil, error
    }

    return tasks, nil
}

func FindTaskById(id string, task *Task) (error)  {
    connection, db, error := DbConnection()

    if error != nil {
        return error
    }

    defer connection.Close()

    var objectId bson.ObjectId

    error = ObjectId(id, &objectId)

    if error != nil {
        return mgo.ErrNotFound
    }

    error = TaskCollection(db).FindId(objectId).One(&task)

    if error != nil {
        return error
    }

    return nil
}

func CreateTask(task Task) error {
    connection, db, error := DbConnection()

    if error != nil {
        return error
    }

    defer connection.Close()

    return TaskCollection(db).Insert(task)
}

func TaskCollection(database *mgo.Database) *mgo.Collection {
    return database.C("task")
}

func DbConnection() (*mgo.Session, *mgo.Database, error) {
    connection, error := mgo.Dial(os.Getenv("MONGODB_URI"))

    if error != nil {
        return nil, nil, error
    }

    return connection, connection.DB(os.Getenv("MONGODB_DB_NAME")), nil
}

func ObjectId(id string, objectId *bson.ObjectId) error {
    if ! bson.IsObjectIdHex(id) {
        return errors.New("Invalid object id")
    }

    *objectId = bson.ObjectIdHex(id)

    return nil
}
