filter = bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 25}}}}
update = bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "updated"}}}}

updateResult, err = collection.UpdateMany(context.TODO(), filter, update)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
