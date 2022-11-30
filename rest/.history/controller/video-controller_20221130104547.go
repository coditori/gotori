package controller

import (
	"rest/models"
	"rest/service"

	"github.com/gin-gonic/gin"
)

// controller interface
type VideoController interface {
	Save(ctx *gin.Context)
	FindAll() []models.Video
}

// contorller class
type videoController struct {
	// injection
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return videoController{
		service: service,
	}
}

func (controller *videoController) Save(ctx *gin.Context) {
	var video models.Video
	ctx.BindJSON(&video)
	controller.service.Save(video)
}

func (controller *videoController) FindAll() []models.Video {
	return controller.FindAll()
}
