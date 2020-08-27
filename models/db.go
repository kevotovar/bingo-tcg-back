package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// db Mongo manager
var db *mongo.Database

// InitDB mongo connection
func InitDB(dbURL string, dbName string) {
	clientOptions := options.Client().ApplyURI(dbURL)
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	db = mongoClient.Database(dbName)
}
