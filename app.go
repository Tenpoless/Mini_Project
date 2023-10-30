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

	os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":8000"))
}
