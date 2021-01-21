package services

import (
	"main/api/types"
	"main/twilio"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

// PrintMessage - prints the message
func PrintMessage(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hey there",
	})
}

// SendMessage - sends a sms, will add different transports later
func SendMessage(ctx *fiber.Ctx) error {
	body := new(types.SendMessage)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	smsStatus := twilio.SendSMS(body.ToNumber, body.Message)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": smsStatus,
	})
}
