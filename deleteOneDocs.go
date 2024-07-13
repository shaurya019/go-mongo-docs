filter = bson.D{{Key: "name", Value: "John Doe"}}

deleteResult, err := collection.DeleteOne(context.TODO(), filter)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Deleted %v documents in the users collection\n", deleteResult.DeletedCount)
