package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func InitMongoDB() {
	// Load environment variables
	godotenv.Load()

	// Get MongoDB URI and Database Name from .env
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to ensure connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Initialize the collection
	TaskCollection = client.Database(dbName).Collection("tasks")
}
