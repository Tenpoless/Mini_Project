package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteStokDarah digunakan untuk menghapus stok darah
func DeleteStokDarah(c echo.Context) error {
    // Pastikan hanya admin yang dapat mengakses endpoint ini
    isAdmin, _ := c.Get("isAdmin").(bool)
    if !isAdmin {
        return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Anda tidak memiliki izin untuk mengakses operasi ini"))
    }

    // Ambil ID stok darah dari URL
    id := c.Param("id")

    // Hapus data stok darah dari database
    if err := config.DB.Where("id = ?", id).Delete(&models.Stok{}).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghapus stok darah"))
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Stok darah berhasil dihapus", nil))
}