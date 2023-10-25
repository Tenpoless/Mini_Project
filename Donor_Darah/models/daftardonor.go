package models

import (
	"gorm.io/gorm"
	"time"
)

type DaftarDonor struct {
	ID        			uint `gorm:"primarykey"`
	CreatedAt 			time.Time
	UpdatedAt			time.Time
	DeletedAt 			gorm.DeletedAt `gorm:"index"`
	ID_User				uint `gorm:"foreignKey:UserID"`
	ID_Jadwal  			uint `gorm:"foreignKey:JadwalID"`
	Waktu_Pendaftaran 	time.Time
	Status              string
}