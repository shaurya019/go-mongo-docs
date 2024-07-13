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
    users := []interface{}{
		bson.D{{Key: "name", Value: "Alice"}, {Key: "age", Value: 25}},
		bson.D{{Key: "name", Value: "Bob"}, {Key: "age", Value: 28}},
	}
	
	insertManyResult, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)	
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
