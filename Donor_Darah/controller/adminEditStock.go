package controller

import (
	"app/models"
	"app/utils"
	"app/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateStokDarah digunakan untuk mengedit stok darah
func UpdateStokDarah(c echo.Context) error {
    // Pastikan hanya admin yang dapat mengakses endpoint ini
    isAdmin, _ := c.Get("isAdmin").(bool)
    if !isAdmin {
        return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Anda tidak memiliki izin untuk mengakses operasi ini"))
    }

    // Ambil ID stok darah dari URL
    id := c.Param("id")

    // Bind request ke objek model StokDarah
    stok := models.Stok{}
    if err := c.Bind(&stok); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    // Update data stok darah di database
    if err := config.DB.Where("id = ?", id).Updates(&stok).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengupdate stok darah"))
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Stok darah berhasil diupdate", stok))
}