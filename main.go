package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdotrezaa/be-learning-week/modules/products"
	"github.com/mdotrezaa/be-learning-week/modules/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	r := gin.Default()

	// Open connection
	dsn := "root:password@tcp(localhost:3306)/learn_week?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logger.Info)),
	})
	if err != nil {
		panic("failed connection")
	}

	productRequestHandler := products.RequestHandler{}
	r.GET("/products", productRequestHandler.GetProducts)
	r.GET("/products/:id", productRequestHandler.GetProductDetail)

	userRequestHandler := users.RequestHandler{
		DB: db,
	}
	r.GET("/users/:id", userRequestHandler.GetUserDetail)

	r.Run()
}
