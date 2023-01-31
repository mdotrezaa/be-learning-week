package songs

import (
	"time"
)

type SongsData struct {
	ID        uint      `json:"id" gorm:"primaryKey";auto_increment;`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	AlbumId   uint      `json:"albumId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type SongsDataInput struct {
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	AlbumId   uint      `json:"albumId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
