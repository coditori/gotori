package server

import (
	"io"
	"os"
	"rest/controller"
	"rest/repository"
	"rest/service"

	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.New()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
)

func Init() {
	// config := config.GetConfig()
	defer videoRepository.CLoseDB()
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
