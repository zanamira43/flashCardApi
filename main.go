package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zanamira43/flashcardApi/database"
	"github.com/zanamira43/flashcardApi/routes"
)

func init() {
	database.Connect()
}
func main() {
	// new fiber instance
	app := fiber.New()

	// cors middleware with frontend
	app.Use(cors.New(cors.Config{
		// AllowCredentials: true,
		AllowOrigins: "*",
		// AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		// AllowHeaders:     "",
	}))

	// intergrate routes endpoint
	routes.Setup(app)

	// api listening port
	app.Listen("0.0.0.0:3001")
}
