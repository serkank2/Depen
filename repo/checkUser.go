package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUser(email string, collection *mongo.Collection) bool {

	result := collection.FindOne(context.TODO(), bson.D{
		{Key: "email", Value: email},
	})
	if result.Err() == mongo.ErrNoDocuments {
		return true
	}

	if result.Err() != nil {
		return false

	}
	return false
}
