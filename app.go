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

	port := ":8000"

	os.Getenv("PORT")

	if err := e.Start(port); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}
