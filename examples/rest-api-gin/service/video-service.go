package service

import (
	"errors"
	"rest/models"
	"rest/repository"
)

type VideoService interface {
	Save(models.Video) (*models.Response, error)
	Update(models.Video) models.Response
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

func (service *videoService) Save(video models.Video) (*models.Response, error) {
	savdVideo, err := service.videoRepository.Save(video)
	if err != nil {
		return nil, err
	}
	return &models.Response{ID: savdVideo.ID}, nil
}

func (service *videoService) Update(video models.Video) models.Response {
	return models.Response{ID: service.videoRepository.Update(video).ID}
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
