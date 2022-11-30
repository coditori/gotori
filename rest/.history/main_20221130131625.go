package main

import (
	"io"
	"os"
	"rest/controller"
	"rest/middlewares"
	"rest/service"

	"github.com/gin-gonic/gin"
	"github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	router := gin.New()

	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), ginDum)

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messsage": "OK!",
		})
	})
	// router.GET("/albums", getAlbums)
	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	router.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	router.Run("localhost:8080")
}

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
