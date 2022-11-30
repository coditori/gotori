package service

type VideoService interface {
	Save(entity.Video) entity.Video
	findAll()
}
