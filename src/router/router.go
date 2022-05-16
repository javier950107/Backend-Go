package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	addressRoutes "api/src/routes/address"
	userRoutes "api/src/routes/user"
)

func SetupRoutes(app *fiber.App) {
	// create the router router/api/
	api := app.Group("/api", logger.New())

	//setup the routes
	userRoutes.SetupUserRoutes(api)
	addressRoutes.SetupAddressRoutes(api)
}
