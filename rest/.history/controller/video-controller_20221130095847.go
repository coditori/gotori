package controller

import "rest/models"

type VideoController interface {
	Save(models.Video) models.Video
	FindAll() []models.Video
}

type videoController struct {
}
