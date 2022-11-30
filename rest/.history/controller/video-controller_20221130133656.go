package controller

import (
	"rest/models"
	"rest/service"

	"github.com/gin-gonic/gin"
)

// controller interface
type VideoController interface {
	Save(ctx *gin.Context) (models.Video, error)
	FindAll() []models.Video
}

// contorller class
type videoController struct {
	// injection
	service service.VideoService
}

// implement the interface
func New(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (controller *videoController) Save(ctx *gin.Context) (models.Video, error) {
	var video models.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return _, err
	}
	controller.service.Save(video)
	return models.Video, nil
}

func (controller *videoController) FindAll() []models.Video {
	return controller.service.FindAll()
}
