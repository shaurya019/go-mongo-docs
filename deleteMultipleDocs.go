filter = bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 25}}}}

deleteResult, err = collection.DeleteMany(context.TODO(), filter)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Deleted %v documents in the users collection\n", deleteResult.DeletedCount)
