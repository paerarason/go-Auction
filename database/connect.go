package database

import (
	"database/sql"
	 _ "github.com/lib/pq"
	 "log"
	 "os"
	 "github.com/joho/godotenv"
)


func DB_connection() {
	connectionString := "mongodb://localhost:27017"
	databaseName := "yourDatabase"
	newCollectionName := "newCollection"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	database := client.Database(databaseName)
    collection := database.Collection(collectionName)
	filter := bson.M{"username": "desiredUsername"}
     
	// Execute the query
	cursor, err := collection.Find(ctx, filter)

	var users []User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
}