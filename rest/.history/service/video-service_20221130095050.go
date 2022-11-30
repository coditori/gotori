package service

import "rest/models"

type VideoService interface {
	Save(models.Video) models.Video
	findAll() []models.Video
}

type videoService struct {
	videos []models.Video
}
