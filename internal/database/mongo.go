package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Client, *mongo.Database, error) {

	err := godotenv.Load(".env")

	if err != nil {
		return nil, nil, fmt.Errorf("error loading .env file: %w", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, nil, fmt.Errorf("mongo connect error: %w", err)
	}

	// Ping the database
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, nil, fmt.Errorf("mongo ping error: %w", err)
	}

	// Select the database
	db := client.Database("footboss")

	fmt.Println("✅ Connected to MONGODB ATLAS")

	return client, db, nil
}