package routes

import (
	"main/api/services"

	"github.com/gofiber/fiber/v2"
)

// MessagesRoutes Initalizes routes for messages
func MessagesRoutes(app fiber.Router) {

	messageRouter := app.Group("/messages")
	messageRouter.Get("/print", services.PrintMessage)
	messageRouter.Post("/send", services.SendMessage)
}
