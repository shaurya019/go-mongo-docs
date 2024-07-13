filter := bson.D{{Key: "name", Value: "John Doe"}}
update := bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: 31}}}}

updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
