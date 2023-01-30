package entity

import (
	"time"
)

type Album struct {
	ID        int64  `gorm:"primaryKey";auto_increment;not_null`
	Name      string `gorm:"type:varchar(100)"`
	Year      uint
	Songs     []Song
	CreatedAt time.Time
	UpdatedAt time.Time
}
