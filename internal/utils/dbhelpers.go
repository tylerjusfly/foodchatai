package utils

import (
	"context"
	"time"

	"github.com/tylerjusfly/foodchatai/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func EmailExists(collection *mongo.Collection, email string) (bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"email": email}
    var result bson.M

    err := collection.FindOne(ctx, filter).Decode(&result)

    if err == mongo.ErrNoDocuments {
        return false, nil // Email does not exist
    } else if err != nil {
        return false, err // Database error
    }

    return true, nil // Email exists
}

func GetUserByEmail(collection *mongo.Collection, email string) (*models.Profile, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"email": email}
    var result models.Profile

    err := collection.FindOne(ctx, filter).Decode(&result)

    if err == mongo.ErrNoDocuments {
        return &result, err
    } else if err != nil {
        return &result, err
    }

    return &result, nil
}