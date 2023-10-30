package config

import (

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

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// var errDB error
	// DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if errDB != nil {
	// 	panic("Gagal terhubung ke database")
	// }

	// fmt.Println("Terhubung ke Database")
}

