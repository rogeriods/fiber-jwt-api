package main

import (
	"log"
	"rogeriods/fiber-jwt-api/configs"
	"rogeriods/fiber-jwt-api/handlers"
	"rogeriods/fiber-jwt-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitDatabase()

	app := fiber.New()

	// Public urls
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Private urls
	protected := app.Group("/api", middlewares.JwtCheckMiddleware)
	protected.Get("/profile", handlers.Profile)

	log.Fatal(app.Listen(":3000"))
}
