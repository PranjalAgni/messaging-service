package main

import (
	"fmt"
	"main/api/routes"
	"main/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "martini",
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	// initalize all routes
	routes.MessagesRoutes(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
