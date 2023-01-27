package entity

import (
	"time"
)

type Album struct {
	ID        uint
	Name      string `gorm:"type:varchar(100)"`
	Year      uint
	Songs     []Song
	CreatedAt time.Time
	UpdatedAt time.Time
}
