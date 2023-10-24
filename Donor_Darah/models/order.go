package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        		uint `gorm:"primarykey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	ID_Pusat 		uint `gorm:"foreignKey:PusatID"`
	ID_User 		uint `gorm:"foreignKey:UserID"`
	ID_GolDarah 	uint `gorm:"foreignKey:Gol_DarahID"`
	Jumlah 			int32
}
