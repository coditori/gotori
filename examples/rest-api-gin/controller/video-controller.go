package controller

import (
	"net/http"
	"rest/models"
	"rest/service"
	"rest/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// controller interface
type VideoController interface {
	Save(ctx *gin.Context)
	FindAll(ctx *gin.Context)
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

func (controller *videoController) Save(ctx *gin.Context) {
	var video models.Video
	if ctx.ShouldBindJSON(&video) != nil && validate.Struct(video) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not bind or validate data!",
		})
	}
	controller.service.Save(video)
	ctx.JSON(http.StatusOK, video)
}

func (controller *videoController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, controller.service.FindAll())
}
