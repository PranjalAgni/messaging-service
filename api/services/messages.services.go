package services

import (
	"fmt"
	"main/api/types"
	"main/core"
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

	done := make(chan bool, 1)

	go core.SendSMS(done, body.ToNumber, body.Message)

	fmt.Println("Okay fine")

	<-done
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "done",
	})
}
