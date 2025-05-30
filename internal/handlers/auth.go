package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tylerjusfly/foodchatai/internal/utils"
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

	user, err := utils.GetUserByEmail(collection, input.Email)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "user email does not exists"})
	}

	go utils.SendTokenEmail(user.Email, "12345")

	return c.Status(200).JSON(fiber.Map{
		"message": "Email sent",
		"id":      user.ID,
	})
}

func VerifyEmail(c *fiber.Ctx) error {

	var input struct {
		Email string `json:"email"`
		Code string `json:"code"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error in Json"})
	}

	if input.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Please provide an email address"})
	}

	if input.Code == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Enter the code sent to your mail"})
	}

		db, ok := c.Locals("db").(*mongo.Database)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection not available"})
	}

	collection := db.Collection("profiles")

	user, err := utils.GetUserByEmail(collection, input.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "user email does not exists"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Email verified",
		"id": user.ID,
	})
}
