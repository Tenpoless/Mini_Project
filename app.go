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

	// port := "3000"

	os.Getenv("PORT")

	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}
