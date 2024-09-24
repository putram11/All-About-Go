package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// connectDB connects to MongoDB and returns the database instance
func connectDB() (*mongo.Database, error) {
	// Get the MongoDB URI from the environment variable
	uri := os.Getenv("DB_URI")
	if uri == "" {
		return nil, fmt.Errorf("DB_URI environment variable is not set")
	}

	// Create a MongoDB client
	clientOptions := options.Client().ApplyURI(uri).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	// Define a timeout for connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB cluster
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB cluster to verify the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected successfully to MongoDB")

	// Select the database
	database := client.Database("project1")

	return database, nil
}

func main() {
	// Call the connectDB function
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// You can now use `db` to interact with the database
	fmt.Println("Database selected:", db.Name())
}
