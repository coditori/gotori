package service

import (
	"rest/models"
	"rest/repository"
)

type VideoService interface {
	Save(models.Video) models.Video
	Update(models.Video)
	Delete(models.Video)
	FindAll() []models.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video models.Video) models.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video models.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video models.Video) {
	service.videoRepository.Save(video)
}

func (service *videoService) FindAll() []models.Video {
	return service.videoRepository.FindAll()
}
