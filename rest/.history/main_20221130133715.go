package main

import (
	"io"
	"net/http"
	"os"
	"rest/controller"
	"rest/middlewares"
	"rest/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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

	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

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
		video, err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, video)
		}
	})

	router.Run("localhost:8080")
}

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }
