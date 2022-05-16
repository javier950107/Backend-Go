package user

import (
	userHandler "api/src/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	// set up route path/api/ <- /user
	app := router.Group("/user")
	//Get user
	app.Get("/", userHandler.GetUsers)
}
