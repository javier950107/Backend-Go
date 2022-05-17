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
	// get the user by id
	app.Get("/:userId", userHandler.GetUserById)
	// create a user in the DB
	app.Post("/", userHandler.CreateUsers)
	// delete a user
	app.Delete("/:userId", userHandler.DeleteUser)
	// update user
	app.Put("/:userId", userHandler.UpdateUser)

}
