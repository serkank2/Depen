package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:1453@localhost:27017/")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

var DB = ConnectionDB()

func GetCollection(client *mongo.Client, collection string) *mongo.Collection {
	return client.Database("depen").Collection(collection)
}
