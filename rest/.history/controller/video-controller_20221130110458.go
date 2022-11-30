package controller

import (
	"rest/models"
	"rest/service"

	"github.com/gin-gonic/gin"
)

// controller interface
type VideoController interface {
	Save(ctx *gin.Context) models.Video
	FindAll() []models.Video
}

// contorller class
type videoController struct {
	// injection
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &videoController{
		service: s,
	}
}

func (controller *videoController) Save(ctx *gin.Context) models.Video {
	var video models.Video
	ctx.BindJSON(&video)
	controller.service.Save(video)
	return video
}

func (controller *videoController) FindAll() []models.Video {
	return controller.service.FindAll()
}
