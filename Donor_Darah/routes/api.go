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
	e.GET("/admin/stok", controller.GetStok)  // Mmengambil informasi stok darah
    e.POST("/admin/stok", controller.CreateStokDarah)  // Menambahkan stok darah baru
    e.PUT("/admin/stok/:id", controller.UpdateStokDarah)  // Memperbarui stok darah
    e.DELETE("/admin/stok/:id", controller.DeleteStokDarah) 
	e.PUT("/admin/daftardonor/:id/status", controller.UpdateStatus)
	e.GET("/admin/daftardonor/:id/status", controller.GetStatus)

	return e
}

