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

	modelTypes := []reflect.Type{
		reflect.TypeOf(models.Admin{}),
		reflect.TypeOf(models.DaftarDonor{}),
		reflect.TypeOf(models.Gol_Darah{}),
		reflect.TypeOf(models.Jadwal{}),
		reflect.TypeOf(models.Order{}),
		reflect.TypeOf(models.Pusat{}),
		reflect.TypeOf(models.Stok{}),
		reflect.TypeOf(models.User{}),
	}
	
	for _, modelType := range modelTypes {
		modelPtr := reflect.New(modelType).Interface()
		config.DB.AutoMigrate(modelPtr)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
