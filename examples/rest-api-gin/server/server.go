package server

import (
	"io"
	"os"
	"rest/controller"
	"rest/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func mains() {
	saveLogToFile()
	router := NewRouter()
	port := applicationPort()
	router.Run(":" + port)
}

func saveLogToFile() {
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func applicationPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "5000"
	}
	return port
}
