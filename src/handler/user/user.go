package userHandler

import (
	"api/src/database"
	"api/src/model"

	"github.com/gofiber/fiber/v2"
)

// get users data func
func GetUsers(response *fiber.Ctx) error {
	db := database.DB       //declare db connection
	var users []model.Users // set up the array instance struct model

	//find data aka: SELECT * FROM users
	db.Find(&users)

	// if the lenght is 0
	if len(users) <= 0 {
		return response.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "No data found",
			"data":   nil})
	}

	// if found data return this data
	return response.Status(200).JSON(fiber.Map{
		"status": "success",
		"msg":    "Data Found",
		"data":   users})
}

// On create a user in the table user
func CreateUsers(response *fiber.Ctx) error {
	//ini data base config
	db := database.DB
	// new object for add
	user := new(model.Users)
	// found the values in the body
	err := response.BodyParser(user)

	// if found a error show it
	if err != nil {
		return response.Status(500).JSON(fiber.Map{
			"status": "error",
			"msg":    "Error body not complete, Could not create user",
			"data":   nil,
		})
	}

	// create user aka: INSERT INTO users bla bla
	err = db.Create(&user).Error

	// if found a error inserted
	if err != nil {
		return response.Status(500).JSON(fiber.Map{
			"status": "error",
			"msg":    "Error on create user",
			"data":   nil,
		})
	}

	return response.JSON(fiber.Map{
		"status": "success",
		"msg":    "User Created",
		"data":   user,
	})
}

// Get a user by id
func GetUserById(response *fiber.Ctx) error {
	db := database.DB // init database
	var user model.Users

	// get the id for search
	id := response.Params("userId") //! This is only a test from param the correct is in the body

	// Find the user by id aka: SELECT * FROM users WHERE id = ?
	err := db.Find(&user, "id = ?", id)

	if err != nil {
		return response.Status(500).JSON(fiber.Map{
			"status": "error",
			"msg":    "Not found user",
			"data":   nil,
		})
	}

	return response.Status(200).JSON(fiber.Map{
		"status": "success",
		"msg":    "Data Found",
		"data":   user,
	})
}

// same as get user
func DeleteUser(response *fiber.Ctx) error {
	db := database.DB // init database
	var user model.Users

	// get the id for search
	id := response.Params("userId") //! This is only a test from param the correct is in the body

	// Find the user by id aka: SELECT * FROM users WHERE id = ?
	db.Find(&user, "id = ?", id)
	// delete a user aka: DELETE FROM user WHERE id = ?
	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return response.Status(500).JSON(fiber.Map{
			"status": "error",
			"msg":    "Error on delete user",
			"data":   nil,
		})
	}

	return response.Status(200).JSON(fiber.Map{
		"status": "success",
		"msg":    "Success deleted",
		"data":   user,
	})
}
