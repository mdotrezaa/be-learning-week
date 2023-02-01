package albums

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

func (h RequestHandler) GetAlbums(c *gin.Context) {
	var album []entity.Album
	h.DB.Find(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
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
		ID:        album.ID,
		Name:      album.Name,
		Songs:     NewSongList(album.Songs),
		Year:      album.Year,
		CreatedAt: album.CreatedAt,
		UpdatedAt: album.UpdatedAt,
	}

	c.JSON(http.StatusOK, dto.Response{Data: albumData})
}

func (h RequestHandler) CreateAlbum(c *gin.Context) {
	var p AlbumsDataInput
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid payload"})
		return
	}

	album := entity.Album{Name: p.Name, Year: p.Year}

	h.DB.Create(&album)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: album})
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
func (h RequestHandler) UpdateAlbum(c *gin.Context) {
	var p entity.Album
	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "Record not found!"})
		return
	}
	c.BindJSON(&p)
	h.DB.Save(&p)
	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: p})
}
