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

func Init() {
	// config := config.GetConfig()
	saveLogToFile()
	router := NewRouter()
	port := getApplicationPort()
	router.Run(":" + port)
}

func saveLogToFile() {
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func getApplicationPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	return port
}
