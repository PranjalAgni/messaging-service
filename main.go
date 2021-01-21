package main

import (
	"fmt"
	"log"
	"main/api/routes"
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
	})

	app.Use(logger.New())

	// initalize all routes
	routes.MessagesRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
