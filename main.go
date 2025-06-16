package main

import (
	"log"
	"rogeriods/fiber-jwt-api/configs"
	"rogeriods/fiber-jwt-api/handlers"
	"rogeriods/fiber-jwt-api/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	configs.InitDatabase()

	app := fiber.New()

	// Enable CORS for all origins
	app.Use(cors.New())

	// Public urls
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Private urls
	protected := app.Group("/api", middlewares.JwtCheckMiddleware)
	protected.Get("/profile", handlers.Profile)

	// Private URLs for contacts
	protected.Post("/contacts", handlers.CreateContact)
	protected.Get("/contacts", handlers.GetContacts)
	protected.Get("/contacts/:id", handlers.GetContactById)
	protected.Put("/contacts/:id", handlers.UpdateContact)
	protected.Delete("/contacts/:id", handlers.DeleteContact)

	log.Fatal(app.Listen(":3000"))
}
