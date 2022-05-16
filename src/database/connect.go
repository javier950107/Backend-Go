package database

import (
	"fmt"
	"log"
	"strconv"

	"api/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/src/config"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error in port")
	}

	// Connection URL database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to DB and initialize

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(fmt.Sprintf("failed to connect DB %v", err))
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&model.Users{}, &model.Address{})
	fmt.Println("Database Migrate")
}
