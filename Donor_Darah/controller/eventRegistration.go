package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegistToEvent(c echo.Context) error {
	//ambil ID_User dan ID_Jadwal
	userID := c.FormValue("ID_User")
	jadwalID := c.FormValue("ID_Jadwal")

	//periksa jadwal 
	jadwal := models.Jadwal{}
	if err := config.DB.Where("ID = ?", jadwalID).First(&jadwal).Error; err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Jadwal donor darah tidak ditemukan"))
    }

	//periksa jadwal penuh
	if jadwal.Kapasitas <= 0 {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Acara donor darah sudah penuh"))
    }

	//ubah userID ke uint
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		// Handle kesalahan konversi, misalnya jika userID bukan angka yang valid
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid user ID"))
	}
	
	//ubah jadwalID ke uint
	jadwalIDInt, err := strconv.Atoi(jadwalID)
	if err != nil {
		// Handle kesalahan konversi, misalnya jika jadwalID bukan angka yang valid
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid jadwal ID"))
	}

	//save pendaftaran
	registration := models.DaftarDonor{
		ID_User: uint(userIDInt),
		ID_Jadwal: uint(jadwalIDInt),
	}

	if err := config.DB.Create(&registration).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mendaftar ke acara donor darah"))
    }

	// Update kapasitas acara donor darah
    jadwal.Kapasitas--

    if err := config.DB.Save(&jadwal).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memperbarui kapasitas acara donor darah"))
    }

	return c.JSON(http.StatusOK, utils.SuccessResponse("Pendaftaran berhasil", nil))
}