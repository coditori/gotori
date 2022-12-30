package service

import (
	"log"
	"rest/models"
	"rest/repository"
)

type VideoService interface {
	Save(models.Video) models.Video
	Update(models.Video)
	Delete(models.Video)
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

func (service *videoService) Delete(video models.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) FindById(id uint64) models.Video {
	var video models.Video
	video.ID = id
	log.Printf("FindById Video is %+v\n", &video)
	return service.videoRepository.FindById(video)
}

func (service *videoService) FindAll() []models.Video {
	return service.videoRepository.FindAll()
}
