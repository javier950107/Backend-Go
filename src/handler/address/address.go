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

// get a join
func GetAddressJoinUser(response *fiber.Ctx) error {
	//create the struct who gives a result
	type Result struct {
		City    string `json:"city"`
		Country string `json:"country"`
	}

	db := database.DB
	// inst the struct
	var result Result
	// make the join
	rows, err := db.Table("users").
		Joins("JOIN addresses ON addresses.id = users.id_address").
		Select("addresses.city, addresses.country").
		Rows()

	// look inside the rows
	for rows.Next() {
		db.ScanRows(rows, &result)
	}

	// validate error
	if err != nil {
		return response.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "Not data found",
			"data":   nil,
		})
	}

	return response.JSON(fiber.Map{
		"status": "success",
		"msg":    "Data success",
		"data":   result,
	})
}
