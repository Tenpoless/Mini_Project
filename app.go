package main

import (
	"app/config"
	"app/routes"
	"log"
)

func main() {
	config.ConnectDB()

	e := routes.Init()

	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}
