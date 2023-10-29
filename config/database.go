package config

import (
	"fmt"
	"os"

	// "path/filepath"

	// "github.com/joho/godotenv"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// err := godotenv.Load(filepath.Join(".", ".env"))
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// 	os.Exit(1)
	// }

	os.Getenv("DB_USER")
	os.Getenv("DB_PASS")
	os.Getenv("DB_HOST")
	os.Getenv("PORT")
	os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	dbUser, dbPass, dbHost, dbPort, dbName)

	// var errDB error
	// DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if errDB != nil {
	// 	panic("Failed to Connect Database")
	// }

	fmt.Println("Connected to Database")
}

