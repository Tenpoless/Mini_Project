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
	e.GET("/user", controller.Index)
	e.GET("/user/:id", controller.Show)
	e.POST("/user/register", controller.Store)
	e.POST("/user/login", controller.Login)
	e.PUT("/user/:id", controller.Update)
	e.DELETE("/user/:id", controller.Delete)
	e.POST("/user/registevent", controller.RegistToEvent)
	e.POST("/user/order", controller.PesanDarah)

	
	//admin routes
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.AuthorizationAdmin)

	adminGroup.POST("/createstok", controller.CreateStokDarah)
	adminGroup.GET("/getstok", controller.GetStok)
	adminGroup.PUT("/updatestok", controller.UpdateStokDarah)
	adminGroup.DELETE("/deletestok", controller.DeleteStokDarah)
	adminGroup.PUT("/daftardonor/:id/updatestatus", controller.UpdateStatus)
	adminGroup.GET("/daftardonor/:id/getstatus", controller.GetStatus)
	
	e.POST("/admin/login", controller.AdminLogin)
	e.GET("/admin/getstok", controller.GetStok)  // Mmengambil informasi stok darah
    e.POST("/admin/createstok", controller.CreateStokDarah)  // Menambahkan stok darah baru
    e.PUT("/admin/updatestok/:id", controller.UpdateStokDarah)  // Memperbarui stok darah
    e.DELETE("/admin/deletestok/:id", controller.DeleteStokDarah) 
	e.PUT("/admin/daftardonor/:id/updatestatus", controller.UpdateStatus)
	e.GET("/admin/daftardonor/:id/getstatus", controller.GetStatus)

	return e
}

