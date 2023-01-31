package entity

import (
	"time"
)

type Album struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Year      uint      `json:"year"`
	Songs     []Song    `json:"songs"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
