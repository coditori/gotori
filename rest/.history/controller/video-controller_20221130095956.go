package controller

import (
	"rest/models"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context)
	FindAll() []models.Video
}

type videoController struct {
}
