package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User mongo struct
type User struct {
	ID    primitive.ObjectID `bson:"id"`
	Email string             `bson:"email"`
}

// Save user data
func (u *User) Save() (*mongo.InsertOneResult, error) {
	collection := db.Collection("users")
	result, err := collection.InsertOne(context.TODO(), u)
	return result, err
}
