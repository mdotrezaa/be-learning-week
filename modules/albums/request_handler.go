package albums

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

func (h RequestHandler) GetAlbums(c *gin.Context) {
	var album []entity.Album
	h.DB.Preload("Album").Find(&album)

	p := make([]AlbumData, len(album))
	for i, album := range album {
		p[i] = AlbumData{
			ID:    album.ID,
			Name:  album.Name,
			Songs: NewSongList(album.Songs),
			Year:  album.Year,
		}
	}

	c.JSON(http.StatusOK, dto.Response{Data: p})
}

func (h RequestHandler) GetAlbumsDetail(c *gin.Context) {
	var album entity.Album
	id := c.Params.ByName("id")
	if err := h.DB.Where("id = ?", id).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "album not found"})
		return
	}

	h.DB.Preload("Songs").First(&album, id)

	albumData := AlbumData{
		ID:    album.ID,
		Name:  album.Name,
		Songs: NewSongList(album.Songs),
		Year:  album.Year,
	}

	c.JSON(http.StatusOK, dto.Response{Data: albumData})
}

func (h RequestHandler) CreateAlbum(c *gin.Context) {
	var p entity.Album
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid payload"})
		return
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	h.DB.Create(p)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: p})
}

func (h RequestHandler) DeleteAlbum(c *gin.Context) {
	var p entity.Album
	id := c.Params.ByName("id")
	if err := h.DB.Where("id = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid id"})
		return
	}

	h.DB.Delete(p, id)

	c.JSON(http.StatusOK, dto.Response{Message: "Album Deleted", Data: true})
}
