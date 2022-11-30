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
type controller struct {
	// injection
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}

func (controller *VideoController) Save(ctx *gin.Context) {

}

func (controller *VideoController) FindAll() []models.Video {

}
