package entity

import (
	"time"
)

type Song struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AlbumId   uint      `json:"albumId"`
	Title     string    `json:"title" gorm:"type:varchar(255)" binding:"required"`
	Author    string    `json:"author" gorm:"type:varchar(100)" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
