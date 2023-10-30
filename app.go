package main

import (
	"app/config"
	"app/routes"
	"log"
	"os"
)

func main() {
	config.ConnectDB()

	e := routes.Init()

	port := "3306"

	os.Getenv("PORT")

	if err := e.Start("0.0.0.0:" +port); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}
