package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Link db struc
type Link struct {
	ID      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"title"`
	Address string             `bson:"address"`
	User    User               `bson:"user"`
}

// Save the link data
func (l *Link) Save() (*mongo.InsertOneResult, error) {
	collection := db.Collection("links")
	insertResult, err := collection.InsertOne(context.TODO(), l)
	return insertResult, err
}

// GetByID retrieves the data
func (l *Link) GetByID(ID string) error {
	collection := db.Collection("links")
	filter := struct{ _id string }{_id: ID}
	err := collection.FindOne(context.TODO(), filter).Decode(&l)
	return err
}
