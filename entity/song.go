package entity

import (
	"time"
)

type Song struct {
	ID        uint `gorm:"primaryKey"`
	AlbumId   uint
	Album     Album
	Title     string `gorm:"type:varchar(255)"`
	Author    string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
