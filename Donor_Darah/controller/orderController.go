package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

//memeriksa stok darah
func GetStokDarah(c echo.Context) error {
    golDarah := c.Param("golDarah")
    
    var darah models.Stok
    if err := config.DB.Where("gol_darah = ?", golDarah).First(&darah).Error; err != nil {
        return c.JSON(http.StatusNotFound, utils.ErrorResponse("Stok darah tidak ditemukan"))
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Stok darah ditemukan", darah))
}


func PesanDarah(c echo.Context) error {
    var pesanan models.Order

    if err := c.Bind(&pesanan); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

	// Mendapatkan data stok darah berdasarkan golongan darah yang diminta
	var stok models.Stok
	if err := config.DB.Where("gol_darah = ?", pesanan.ID_GolDarah).First(&stok).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Stok darah tidak ditemukan"))
	}

	//cek stok
	if pesanan.Jumlah > int32(stok.Jumlah) {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Stok darah tidak mencukupi untuk pesanan ini"))
	}

	//jika stok mencukupi untuk pesanan
	if pesanan.Jumlah <= int32(stok.Jumlah) {
		// Simpan pesanan ke dalam database
		if err := config.DB.Create(&pesanan).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menyimpan pesanan darah"))
		}
		
		// Update stok
		stok.Jumlah -= int64(pesanan.Jumlah)
		if err := config.DB.Save(&stok).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memperbarui stok darah"))
		}
	
		return c.JSON(http.StatusOK, utils.SuccessResponse("Pesanan darah berhasil", pesanan))
	}

    return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Pesanan tidak berhasil"))
}
