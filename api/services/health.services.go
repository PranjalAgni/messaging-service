package services

import "github.com/gofiber/fiber/v2"

func HealthMonitor(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"health": "Running",
	})
}
