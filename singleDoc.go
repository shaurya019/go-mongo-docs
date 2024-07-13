package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

func main() {
    // Setup connection (same as above)
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    // Get a handle for your collection
    collection := client.Database("testdb").Collection("users")

    // Create a BSON document
    user := bson.D{
        {Key: "name", Value: "John Doe"},
        {Key: "age", Value: 30},
        {Key: "address", Value: bson.D{
            {Key: "street", Value: "123 Main St"},
            {Key: "city", Value: "New York"},
        }},
    }

    // Insert a single document
    insertResult, err := collection.InsertOne(context.TODO(), user)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
