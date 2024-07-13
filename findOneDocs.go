var result bson.M
err = collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: "John Doe"}}).Decode(&result)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found a single document: %+v\n", result)
