package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// PrintMessage - prints the message
func PrintMessage(c *fiber.Ctx) error {
	fmt.Println("Inside here")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hey there",
	})
}
