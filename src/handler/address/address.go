package addressHandler

import (
	"api/src/database"
	"api/src/model"

	"github.com/gofiber/fiber/v2"
)

// Get address info
func GetAddress(response *fiber.Ctx) error {
	db := database.DB           //init the DB interaction
	var address []model.Address // create a empty array reference struct address

	//find all data aka: SELECT * FROM address
	db.Find(&address)

	// if not found data validation
	if len(address) <= 0 {
		return response.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "Not data found",
			"data":   nil})
	}

	// if found data send
	return response.JSON(fiber.Map{
		"status": "success",
		"msg":    "Data found",
		"data":   address})
}
