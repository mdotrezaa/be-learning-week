package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdotrezaa/be-learning-week/dto"
)

type RequestHandler struct{}

func (RequestHandler) GetProducts(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid id"})
		return
	}
	p := Songs{
		ID:     uint(id),
		Title:  "Toyota",
		Author: "Toyota",
	}
	c.JSON(http.StatusOK, dto.Response{Data: p})
}

func (RequestHandler) GetProductDetail(c *gin.Context) {
	prds := make([]Songs, 0)
	c.JSON(http.StatusOK, dto.Response{Data: prds})
}
