package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	addressRoutes "api/src/routes/address"
	userRoutes "api/src/routes/user"
)

func SetupRoutes(app *fiber.App) {
	// create the router router/api/
	// read the middleware
	api := app.Group("/api", func(response *fiber.Ctx) error {
		fmt.Println("Read the middleware")
		return response.Next()
	})

	//setup the routes
	userRoutes.SetupUserRoutes(api)
	addressRoutes.SetupAddressRoutes(api)
}
