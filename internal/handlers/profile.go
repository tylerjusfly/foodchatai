package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tylerjusfly/foodchatai/internal/models"
	"github.com/tylerjusfly/foodchatai/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProfile(c *fiber.Ctx) error {

	idParam := c.Params("id")

	if idParam == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	// Convert string to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	// Get database from context
	db, ok := c.Locals("db").(*mongo.Database)

	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database connection not available",
		})
	}

	collection := db.Collection("profiles")

	filter := bson.M{"_id": objectID}

	var result models.Profile

	err = collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(200).JSON(fiber.Map{
				"result": nil,
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to query database",
		})
	}

	return c.Render("profile", fiber.Map{
			"Name": result.Name,
			"Email": result.Email,
			"ID": result.ID.Hex(),
		})
}

func CreateProfileView(c *fiber.Ctx) error{

	return c.Render("createprofile", fiber.Map{
			"Title": fmt.Sprintf("Hello! %s Welcome to FootBoss", "m"),
		})

}

func CreateProfile(c *fiber.Ctx) error{
	
	var input struct {
		Name    string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&input); err != nil{
		return c.Status(400).JSON(fiber.Map{"error": "error in Json"})
	}

	if input.Name == ""{
		return c.Status(400).JSON(fiber.Map{"error": "name is required"})
	}

	if input.Email == ""{
		return c.Status(400).JSON(fiber.Map{"error": "email is required"})
	}

	db, ok := c.Locals("db").(*mongo.Database)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection not available"})
	}

	collection := db.Collection("profiles")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

		// Check for email
	userExists, err := utils.EmailExists(collection, input.Email)

	if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Something went wrong",
        })
    }

    if userExists {
        return c.Status(400).JSON(fiber.Map{
            "error": "Email already in use",
        })
    }

	owner := models.Profile{
		Name: input.Name,
		Email: input.Email,
		AiKey: "",
		AiType: "",
	}

	result, err := collection.InsertOne(ctx, owner)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to insert profile"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Club owner created successfully",
		"id":      result.InsertedID,
	})

}