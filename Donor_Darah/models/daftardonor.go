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
	ID_User				uint `gorm:"foreignKey:ID_User"`
	ID_Jadwal  			uint `gorm:"foreignKey:ID_Jadwal"`
	Waktu_Pendaftaran 	time.Time
	Status              string
}

func (u *DaftarDonor) TableName() string {
    return "DaftarDonor"
}