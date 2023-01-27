package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		products, _ := getProducts()
		c.JSON(http.StatusOK, response{Data: products})
	})

	r.GET("/products/:id", func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.ParseInt(idstr, 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, response{Message: "invalid id"})
			return
		}
		p := product{
			ID:    uint(id),
			Name:  "Toyota",
			Price: 200000000,
		}
		c.JSON(http.StatusOK, response{Data: p})
	})

	r.POST("/products", func(c *gin.Context) {
		var p product
		if err := c.Bind(&p); err != nil {
			c.JSON(http.StatusBadRequest, response{Message: "invalid payload"})
			return
		}
		p.ID = 1
		p.CreatedAt = time.Now()
		p.UpdatedAt = time.Now()
		c.JSON(http.StatusOK, response{Message: "success", Data: p})
	})

	r.Run()
}

func getProducts() ([]product, error) {
	products := []product{
		{
			ID:        1,
			Name:      "Honda",
			Price:     100000000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return products, nil
}
