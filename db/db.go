package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client


// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://sooa_mongo_admin:CbfRdzY1dULYKIiE@sooa-mongo-cluster.lrlq0px.mongodb.net/?retryWrites=true&w=majority")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	fmt.Println("⛁ Connected to Database")

	

	return client
}
