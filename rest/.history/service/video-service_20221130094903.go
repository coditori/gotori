package service

type VideoService interface {
	Save(models.Video) models.Vided
	findAll()
}
