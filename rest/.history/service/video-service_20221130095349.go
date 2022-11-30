package service

import "rest/models"

type VideoService interface {
	Save(models.Video) models.Video
	findAll() []models.Video
}

type videoService struct {
	videos []models.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(models.Video) models.Video {

}

func (service *videoService) FinAll() []models.Video {

}
