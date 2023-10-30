package main

import (
	"app/config"
	"app/routes"
	// "log"
	"os"
)

func main() {
	config.ConnectDB()

	e := routes.Init()

	port := "8000"

	os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" +port))
}
