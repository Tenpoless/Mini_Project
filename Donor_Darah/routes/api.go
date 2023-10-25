package routes

import (
	"app/controller"
	"app/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ayo Donor Darah!!!")
	})

	//user routes
	e.GET("/users", controller.Index)
	e.GET("/users/:id", controller.Show)
	e.POST("/users/register", controller.Store)
	e.POST("/users/login", controller.Login)
	e.PUT("/users/:id", controller.Update)
	e.DELETE("/users/:id", controller.Delete)

	//admin routes
	e.POST("/admin/login", controller.AdminLogin)

	//group routes admin 
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.(controller.AdminJWTConfig))

	adminGroup.GET("/stok", controller.GetStok)  // Misalnya, mengambil informasi stok darah
    adminGroup.POST("/stok", controller.CreateStokDarah)  // Menambahkan stok darah baru
    adminGroup.PUT("/stok/:id", controller.UpdateStokDarah)  // Memperbarui stok darah
    adminGroup.DELETE("/stok/:id", controller.DeleteStokDarah) 

	return e

}

