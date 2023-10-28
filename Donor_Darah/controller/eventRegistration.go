package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func RegistToEvent(c echo.Context) error {
	// Ambil Nama dan ID_Jadwal dari JSON request
    nama := c.FormValue("Nama")
    jadwalIDStr := c.FormValue("ID_Jadwal")

    // Konversi ID_Jadwal ke uint
    jadwalID, err := strconv.ParseUint(jadwalIDStr, 10, 0)
    if err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID_Jadwal"))
    }

    // Periksa apakah jadwal dengan ID_Jadwal ada dalam database
    jadwal := models.Jadwal{}
    if err := config.DB.Where("id = ?", jadwalID).First(&jadwal).Error; err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Jadwal donor darah tidak ditemukan"))
    }

    // Periksa apakah jadwal penuh
    if jadwal.Kapasitas <= 0 {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Acara donor darah sudah penuh"))
    }

    // Simpan pendaftaran
    registration := models.DaftarDonor{
        Name:              nama,
        ID_Jadwal:         uint(jadwalID),
        Waktu_Pendaftaran: time.Now(),             // Waktu sekarang
        Status:            "Pendaftaran Berhasil", // Atur status pendaftaran
    }

    if err := config.DB.Create(&registration).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mendaftar ke acara donor darah"))
    }

    // Perbarui kapasitas acara donor darah
    jadwal.Kapasitas--

    if err := config.DB.Save(&jadwal).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memperbarui kapasitas acara donor darah"))
    }

    // Kembalikan respons dengan data yang diminta
    response := struct {
        ID                uint      `json:"ID"`
        Name              string    `json:"Name"`
        ID_Jadwal         uint      `json:"ID_Jadwal"`
        Waktu_Pendaftaran time.Time `json:"Waktu_Pendaftaran"`
        Status            string    `json:"Status"`
    }{
        ID:                registration.ID,
        Name:              registration.Name,
        ID_Jadwal:         registration.ID_Jadwal,
        Waktu_Pendaftaran: registration.Waktu_Pendaftaran,
        Status:            registration.Status,
    }

	// response := web.userRegistResponse{
	// 	Email:    user.Email,
	// 	Password: loginRequest.Password,
	// 	Token:    token,
	// }

	// return c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", response))


    return c.JSON(http.StatusOK, utils.SuccessResponse("Pendaftaran berhasil", response))
}
