package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/tylerjusfly/foodchatai/internal/database"
	"github.com/tylerjusfly/foodchatai/internal/handlers"
)

func main() {

	client, db, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.TODO())

	// Example: list collections
	collections, err := db.ListCollectionNames(context.TODO(), map[string]interface{}{})
	if err != nil {
		log.Fatalf("List collections failed: %v", err)
	}

	fmt.Println("Collections in 'footboss':", collections)

	 // Initialize a standard Go's html/template engine
  	engine := html.New("./views", ".html")
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		AppName: "Footboss",
		Views: engine,
	})

	// Pass the db instance to your handlers using Fiber's context
    app.Use(func(c *fiber.Ctx) error {
        c.Locals("db", db)  // Store db in context
        return c.Next()
    })

// 	 app.Use(limiter.New(limiter.Config{
//       Expiration: 10 * time.Second,
//       Max:        3,
//   }))

	app.Get("/", handlers.GetToHomePage)
	app.Get("/login", handlers.GetToLoginPage)
	app.Post("/login", handlers.SignIn)
	app.Post("/verify", handlers.VerifyEmail)

	// Profile Routes
	app.Get("/profile/create", handlers.CreateProfileView)
	app.Get("/profile/:id", handlers.GetProfile)
	
	app.Post("/api/profile", handlers.CreateProfile)

	// Chat Routes
	app.Get("/chat", handlers.GetChatIndex)
	app.Post("/chat", handlers.CreateChat)
	app.Get("/chat/:chatId", handlers.GetChatPage)
	
	app.Post("/ai/reply", handlers.ReplyChat)

	app.Use(func(c *fiber.Ctx) error {
		// return c.Status(400)
		return c.Render("error", fiber.Map{
			"Route": "/",
		})
	})

	// Get PORT From Env
	PORT := os.Getenv("PORT")

	if PORT == ""{
		PORT = "4000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}