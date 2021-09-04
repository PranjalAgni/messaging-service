package routes

import (
	"main/api/services"

	"github.com/gofiber/fiber/v2"
)

func HealthRoutes(app fiber.Router) {
	app.Get("/", services.HealthMonitor)
}
