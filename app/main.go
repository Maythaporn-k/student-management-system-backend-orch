package main

import (
	"orch/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()

	// *Set content type
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://maythaporn-k.github.io/student-management-system-fronted/home",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Routes
	app.Get("/", indexHandler)
	app.Get("/orch/student-list", handlers.StudentList)
	app.Post("/orch/create-user", handlers.CreateUser)
	app.Delete("/orch/delete-user", handlers.DeleteUser)
	app.Put("/orch/edit-user", handlers.EditUser)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}
	app.Listen(":" + port)
}
