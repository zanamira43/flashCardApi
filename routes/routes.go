package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/flashcardApi/controllers"
)

func Setup(app *fiber.App) {
	// home function endpoint
	app.Get("/", controllers.Home)

	// all deck routes
	app.Get("/api/decks", controllers.GetALLDecks)
	app.Post("/api/decks", controllers.CreateNewDeck)
	app.Get("/api/decks/:id", controllers.GetDeck)
	app.Put("/api/decks/:id", controllers.UpdateDeck)
	app.Delete("/api/decks/:id", controllers.DeleteDeck)

	// all card routes
	app.Get("/api/cards", controllers.GetALLCards)
	app.Post("/api/cards", controllers.CreateNewCard)
	app.Get("/api/cards/:id", controllers.GetCard)
	app.Put("/api/cards/:id", controllers.UpdateCard)
	app.Delete("/api/cards/:id", controllers.DeleteCard)

}
