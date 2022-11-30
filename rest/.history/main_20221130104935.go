package main

import (
	"rest/controller"
	"rest/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	router := gin.Default()
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messsage": "OK!",
		})
	})
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	router.POST("/posts", func(ctx *gin.Context) {

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
