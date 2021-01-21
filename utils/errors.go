package utils

import "github.com/gofiber/fiber/v2"

type httpError struct {
	Statuscode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&httpError{
		Statuscode: code,
		Message:    err.Error(),
	})
}
