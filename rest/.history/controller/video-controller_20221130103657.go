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
	servie service.VideoService
}

func New(s service.VideoService) VideoController {
	return videoController{
		service: s,
	}
}
