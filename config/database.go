package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	os.Getenv("DB_USER")
	os.Getenv("DB_PASS")
	os.Getenv("DB_HOST")
	os.Getenv("DB_PORT")
	os.Getenv("DB_NAME")

	fmt.Println("Connected to Database")
}

