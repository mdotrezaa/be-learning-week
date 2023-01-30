package entity

import (
	"time"
)

type Song struct {
	ID        uint `gorm:"primaryKey";auto_increment;not_null`
	AlbumId   uint
	Album     Album
	Title     string `gorm:"type:varchar(255)"`
	Author    string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
