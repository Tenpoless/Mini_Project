package models

import (
	"gorm.io/gorm"
	"time"
)

type Jadwal struct {
	ID        	uint `gorm:"primarykey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	ID_Pusat	uint `gorm:"foreignKey:PusatID"`
	Tanggal		time.Time
	Kapasitas   int32
}
