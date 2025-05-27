package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tylerjusfly/foodchatai/internal/generators"
	"github.com/tylerjusfly/foodchatai/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetChatIndex(c *fiber.Ctx) error {

	var userId = c.Query("userId")

	if userId == "" {
		return c.RedirectBack("/")
	}

	return c.Render("startchat", fiber.Map{
		"UserID": userId,
	})
}

func GetChatPage(c *fiber.Ctx) error {

	var chatid = c.Params("chatId")
	// var userId = c.Query("userId")

	// Convert string to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(chatid)
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

	collection := db.Collection("chats")

	filter := bson.M{"_id": objectID}

	var result models.Chat

	err = collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(200).JSON(fiber.Map{
				"result": nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query database",
		})
	}

	return c.Render("chat", fiber.Map{
		"UserID": result.UserID.Hex(),
	})
}

func CreateChat(c *fiber.Ctx) error {
	// get userID
	var userId = c.Query("userId")

	if userId == "" {
		return c.Status(400).JSON(fiber.Map{"error": "userId is required"})
	}

	objectID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UserId format"})
	}

	// Create chat
	chat := models.Chat{
		UserID: objectID,
		Title:  "food advice",
	}

	db, ok := c.Locals("db").(*mongo.Database)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection not available"})
	}

	collection := db.Collection("chats")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, chat)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to insert chat"})
	}

	// return id
	return c.Status(201).JSON(fiber.Map{
		"id": result.InsertedID,
	})
}

func ReplyChat(c *fiber.Ctx) error {

	var input struct {
		Text    string `json:"text"`
		Chatid string `json:"chat_id"`
	}

	if err := c.BodyParser(&input); err != nil{
		return c.Status(400).JSON(fiber.Map{"error": "error in Json"})
	}

	if input.Text == ""{
		return c.Status(400).JSON(fiber.Map{"error": "text cannot be empty"})
	}

	if input.Chatid == ""{
		return c.Status(400).JSON(fiber.Map{"error": "chat id is required"})
	}

	answers, err := generators.OpenaiGenerator(input.Text)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "failed to connect to openai"})
	}

	// Save In redis for 1hr

	return c.Status(201).JSON(fiber.Map{
		"response": answers,
	})

}
