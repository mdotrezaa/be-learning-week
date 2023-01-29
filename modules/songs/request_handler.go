package songs

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdotrezaa/be-learning-week/dto"
	"github.com/mdotrezaa/be-learning-week/entity"
	"gorm.io/gorm"
)

type RequestHandler struct {
	DB *gorm.DB
}

func (h RequestHandler) GetSongs(c *gin.Context) {
	var songs []entity.Song
	h.DB.Preload("Songs").Find(&songs)

	p := make([]SongsData, len(songs))
	for i, song := range songs {
		p[i] = SongsData{
			ID:      song.ID,
			Title:   song.Title,
			AlbumId: song.AlbumId,
			Author:  song.Author,
		}
	}

	c.JSON(http.StatusOK, dto.Response{Data: p})
}

func (h RequestHandler) GetSongDetail(c *gin.Context) {
	var songs []entity.Song
	id := c.Params.ByName("id")
	if err := h.DB.Where("album_id = ?", id).First(&songs).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "album not found"})
		return
	}

	h.DB.Where("album_id = ?", id).Find(&songs)
	p := make([]SongsData, len(songs))

	for i, song := range songs {
		p[i] = SongsData{
			ID:      song.ID,
			Title:   song.Title,
			AlbumId: song.AlbumId,
			Author:  song.Author,
		}
	}

	c.JSON(http.StatusOK, dto.Response{Data: p})
}

func (h RequestHandler) CreateSong(c *gin.Context) {
	var p entity.Song
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid payload"})
		return
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	h.DB.Create(p)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: p})
}
func (h RequestHandler) UpdateSong(c *gin.Context) {
	var p entity.Song
	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "data not found"})
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
