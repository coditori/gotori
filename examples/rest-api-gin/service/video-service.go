package service

import (
	"errors"
	"rest/models"
	"rest/repository"
)

type VideoService interface {
	Save(models.Video) models.Video
	Update(models.Video)
	Delete(id uint64) (*models.Video, error)
	FindById(id uint64) models.Video
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

func (service *videoService) Delete(id uint64) (*models.Video, error) {
	video := service.FindById(id)
	if video == (models.Video{}) {
		return nil, errors.New("could not find the resurce")
	}
	service.videoRepository.Delete(video)
	return &video, nil
}

func (service *videoService) FindById(id uint64) models.Video {
	return service.videoRepository.FindById(id)
}

func (service *videoService) FindAll() []models.Video {
	return service.videoRepository.FindAll()
}
