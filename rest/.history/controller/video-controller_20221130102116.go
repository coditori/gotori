package controller

import (
	"rest/models"
	"rest/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context)
	FindAll() []models.Video
}

type videoController struct {
	servie service.VideoService
}
