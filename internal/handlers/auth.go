package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tylerjusfly/foodchatai/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignIn(c *fiber.Ctx) error {

	var input struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error in Json"})
	}

	if input.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "email is required"})
	}

	db, ok := c.Locals("db").(*mongo.Database)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection not available"})
	}

	collection := db.Collection("profiles")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": input.Email}
	var result models.Profile

	err := collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(200).JSON(fiber.Map{
				"error": "This user is not registered yet",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to query database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Email sent",
		"id": result.ID.Hex(),
	})
}
func VerifyEmail(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"message": "Email verified",
	})
}
