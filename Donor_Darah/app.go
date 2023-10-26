package main

import (
	"app/config"
	"app/models"
	"app/routes"
	"reflect"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	models := []interface{}{
		reflect.TypeOf(models.Admin{}),
		reflect.TypeOf(models.DaftarDonor{}),
		reflect.TypeOf(models.Gol_Darah{}),
		reflect.TypeOf(models.Jadwal{}),
		reflect.TypeOf(models.Order{}),
		reflect.TypeOf(models.Pusat{}),
		reflect.TypeOf(models.Stok{}),
		reflect.TypeOf(models.User{}),
	}
	
	for _, m := range models {
		config.DB.AutoMigrate(m)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
