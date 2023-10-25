package controller

import (
	"app/models"
	"app/config"
	"app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateStatus(c echo.Context) error {
    // Ambil ID peserta dari URL parameter
    registrationID := c.Param("id")

    // Ambil data status yang dikirim oleh admin dalam request body
    var statusUpdate models.DaftarDonor
    if err := c.Bind(&statusUpdate); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    // Perbarui status peserta di basis data
    var registration models.DaftarDonor
    if err := config.DB.Where("id = ?", registrationID).First(&registration).Error; err != nil {
        return c.JSON(http.StatusNotFound, utils.ErrorResponse("Registration not found"))
    }

    registration.Status = statusUpdate.Status
    if err := config.DB.Save(&registration).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update registration status"))
    }

    // Beri respons yang sesuai
    return c.JSON(http.StatusOK, utils.SuccessResponse("Registration status updated successfully", registration))
}

// GetStatus memungkinkan admin untuk mengambil status peserta berdasarkan ID
func GetStatus(c echo.Context) error {
    // Ambil ID peserta dari URL parameter
    registrationID := c.Param("id")

    // Ambil status peserta dari basis data berdasarkan ID
    var registration models.DaftarDonor
    if err := config.DB.Where("id = ?", registrationID).First(&registration).Error; err != nil {
        return c.JSON(http.StatusNotFound, utils.ErrorResponse("Registration not found"))
    }

    // Beri respons yang sesuai
    return c.JSON(http.StatusOK, utils.SuccessResponse("Registration status retrieved successfully", registration))
}

// CreateRegistration memungkinkan pengguna untuk membuat pendaftaran peserta donor darah
func CreateRegistration(c echo.Context) error {
    // Ambil data pendaftaran dari request body
    var registration models.DaftarDonor
    if err := c.Bind(&registration); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    // Simpan data pendaftaran ke basis data
    if err := config.DB.Create(&registration).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create registration"))
    }

    // Beri respons yang sesuai
    return c.JSON(http.StatusCreated, utils.SuccessResponse("Registration created successfully", registration))
}