package main

import (
	"fmt"
	"log"
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
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "martini",
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	// initalize all routes
	routes.HealthRoutes(app)
	routes.MessagesRoutes(app)

	log.Println("Server running on http://localhost:" + fmt.Sprintf("%v", os.Getenv("APP_PORT")))
	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))

}
