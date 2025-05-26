package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetToHomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title1": "FoodChat",
		"Title2": "AI",
	})
}

func GetToLoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}