package main

import (
	"app/config"
	"app/routes"
	"os"
)

func main() {
	config.ConnectDB()

	e := routes.Init()

	os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":8000"))
}
