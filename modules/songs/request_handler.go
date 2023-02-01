package songs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdotrezaa/be-learning-week/dto"
	"github.com/mdotrezaa/be-learning-week/entity"
	"gorm.io/gorm"
)

type RequestHandler struct {
	DB *gorm.DB
}

func (h RequestHandler) GetSongs(c *gin.Context) {
	var song []entity.Song
	h.DB.Find(&song)

	c.JSON(http.StatusOK, dto.Response{Data: song, Message: "success"})
}

func (h RequestHandler) GetSongDetail(c *gin.Context) {
	var songs []entity.Song
	id := c.Params.ByName("id")
	if err := h.DB.Where("album_id = ?", id).First(&songs).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "song not found"})
		return
	}

	h.DB.Where("album_id = ?", id).Find(&songs)
	p := make([]SongsData, len(songs))

	for i, song := range songs {
		p[i] = SongsData{
			ID:        song.ID,
			Title:     song.Title,
			AlbumId:   song.AlbumId,
			Author:    song.Author,
			CreatedAt: song.CreatedAt,
			UpdatedAt: song.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: p})
}

func (h RequestHandler) CreateSong(c *gin.Context) {
	var p SongsDataInput
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid payload"})
		return
	}

	song := entity.Song{Title: p.Title, Author: p.Author, AlbumId: p.AlbumId}

	h.DB.Create(&song)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: song})
}

func (h RequestHandler) UpdateSong(c *gin.Context) {
	var p entity.Song
	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "Record not found"})
		return
	}
	c.BindJSON(&p)
	h.DB.Save(&p)
	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: p})
}

func (h RequestHandler) DeleteSong(c *gin.Context) {
	var p entity.Song
	id := c.Params.ByName("id")
	if err := h.DB.Where("id = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "Song not found"})
		return
	}

	h.DB.Delete(p, id)

	c.JSON(http.StatusOK, dto.Response{Message: "Song Deleted", Data: true})
}
