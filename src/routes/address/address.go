package address

import (
	"github.com/gofiber/fiber/v2"

	addressHandler "api/src/handler/address"
)

func SetupAddressRoutes(router fiber.Router) {
	app := router.Group("/address")
	// get address data
	app.Get("/", addressHandler.GetAddress)
	app.Get("/join", addressHandler.GetAddressJoinUser)
}
