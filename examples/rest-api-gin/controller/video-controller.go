package controller

import (
	"log"
	"net/http"
	"rest/models"
	"rest/service"
	"rest/validators"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// controller interface
type VideoController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

func validateRequestAndShowErrorIfCouldNotValidate(ctx *gin.Context, video *models.Video) {
	if ctx.ShouldBindJSON(&video) != nil && validate.Struct(video) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not bind or validate data!",
		})
	}
}

func (controller *videoController) Save(ctx *gin.Context) {
	var video models.Video
	validateRequestAndShowErrorIfCouldNotValidate(ctx, &video)
	log.Printf("Video is %s /n", video)
	controller.service.Save(video)
	ctx.JSON(http.StatusOK, video)
}

func (controller *videoController) Update(ctx *gin.Context) {
	var video models.Video
	validateRequestAndShowErrorIfCouldNotValidate(ctx, &video)

	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 0, 0)
	if err != nil {
		log.Println("Coud not parse param!")
	}
	video.ID = id
	controller.service.Update(video)
}

func (controller *videoController) Delete(ctx *gin.Context) {
	var video models.Video
	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 0, 0)
	if err != nil {
		log.Println("Coud not parse param!")
	}
	video.ID = id
	controller.service.Delete(video)
}

func (controller *videoController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, controller.service.FindAll())
}
