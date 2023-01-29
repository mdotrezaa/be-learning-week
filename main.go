package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mdotrezaa/be-learning-week/modules/albums"
	"github.com/mdotrezaa/be-learning-week/modules/songs"

	"github.com/mdotrezaa/be-learning-week/entity"
)

func main() {
	r := gin.Default()

	// Open connection
	dsn := "root@tcp(localhost:3306)/learn_week?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logger.Info)),
	})
	if err != nil {
		panic("failed connection")
	}
	var album entity.Album
	var song entity.Song
	err = db.AutoMigrate(album, song)
	if err != nil {
		log.Println(fmt.Errorf("db.AutoMigrate: %w", err))
		return
	}

	productRequestHandler := songs.RequestHandler{
		DB: db,
	}
	r.GET("/songs", productRequestHandler.GetSongs)          //Get All Song
	r.GET("/songs/:id", productRequestHandler.GetSongDetail) // Get specific song by album id
	r.POST("/songs", productRequestHandler.CreateSong)       // Add song
	r.PUT("/songs/:id", productRequestHandler.UpdateSong)    // Add song
	r.DELETE("/songs/:id", productRequestHandler.DeleteSong) // Delete song

	userRequestHandler := albums.RequestHandler{
		DB: db,
	}
	r.GET("/albums", userRequestHandler.GetAlbums)           //Get All Album
	r.GET("/albums/:id", userRequestHandler.GetAlbumsDetail) // Get Album Detail
	r.POST("/albums", userRequestHandler.CreateAlbum)        // Add Album
	r.DELETE("/albums/:id", userRequestHandler.DeleteAlbum)  // Add Album

	r.Run()
}
