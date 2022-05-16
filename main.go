package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//start the fiber
	app := fiber.New()
	// database connect DB
	database.ConnectDB()

	// detect routes
	router.SetupRoutes(app)

	port := fmt.Sprintf(":%s", config.Config("PORT_LISTEN"))
	//declare port for listen
	app.Listen(port)
}
