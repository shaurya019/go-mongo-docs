cur, err := collection.Find(context.TODO(), bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 25}}}})
if err != nil {
    log.Fatal(err)
}
defer cur.Close(context.TODO())

for cur.Next(context.TODO()) {
    var elem bson.M
    err := cur.Decode(&elem)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found document: %+v\n", elem)
}

if err := cur.Err(); err != nil {
    log.Fatal(err)
}
