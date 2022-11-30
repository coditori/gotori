package controller

import (
	"rest/models"
	"rest/service"
	"rest/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// controller interface
type VideoController interface {
	Save(ctx *gin.Context) (models.Video, error)
	FindAll() []models.Video
}

// contorller class
type videoController struct {
	// injection
	service service.VideoService
}

var validate *validator.Validate

// implement the interface
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (controller *videoController) Save(ctx *gin.Context) (models.Video, error) {
	var video models.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return video, err
	}
	err = validate.Struct(video)
	if err != nil {
		return video, err
	}
	controller.service.Save(video)
	return video, nil
}

func (controller *videoController) FindAll() []models.Video {
	return controller.service.FindAll()
}
