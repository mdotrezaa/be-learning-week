package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdotrezaa/be-learning-week/dto"
	"github.com/mdotrezaa/be-learning-week/entity"
	"gorm.io/gorm"
)

type RequestHandler struct {
	DB *gorm.DB
}

func (h RequestHandler) GetUserDetail(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid id"})
		return
	}

	var album entity.Album
	h.DB.Preload("Songs").First(&album, id)

	albumData := AlbumData{
		ID:    album.ID,
		Name:  album.Name,
		Songs: NewSongList(album.Songs),
		Year:  album.Year,
	}

	c.JSON(http.StatusOK, dto.Response{Data: albumData})
}
