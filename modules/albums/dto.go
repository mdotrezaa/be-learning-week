package albums

import (
	"time"

	"github.com/mdotrezaa/be-learning-week/entity"
)

type AlbumsDataInput struct {
	Name      string    `json:"name"`
	Year      uint      `json:"year"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type AlbumData struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Songs     []SongData `json:"songs"`
	Year      uint       `json:"year"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type SongData struct {
	ID        uint      `json:"id"`
	AlbumId   uint      `json:"albumId"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewSongList(songs []entity.Song) []SongData {
	songData := make([]SongData, len(songs))
	for i, order := range songs {
		songData[i] = SongData{
			ID:      order.ID,
			AlbumId: order.AlbumId,
			Title:   order.Title,
			Author:  order.Author,
		}
	}

	return songData
}
