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
	FindById(ctx *gin.Context)
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

// PostVideo     godoc
// @Summary      Save a new video
// @Description  Takes a video JSON and save it in DB. Return saved JSON.
// @Tags         videos
// @Produce      json
// @Param        video  body      models.Video  true  "Video JSON"
// @Success      200    {object}  models.Video
// @Router       /auth/videos [post]
// @Security     JWT
func (controller *videoController) Save(ctx *gin.Context) {
	var video models.Video
	validateRequestAndShowErrorIfCouldNotValidate(ctx, &video)
	// log.Printf("Video is %s /n", video)
	controller.service.Save(video)
	ctx.JSON(http.StatusOK, video)
}

// PutVideo      godoc
// @Summary      Update an existing video
// @Description  Takes a video JSON and video id and update it in DB. Return saved JSON.
// @Tags         videos
// @Produce      json
// @Param        uint   path      uint        true  "VideoId"
// @Param        video  body      models.Video  true  "Video JSON"
// @Success      200    {object}  models.Video
// @Router       /auth/videos/{id} [put]
// @Security     JWT
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

// DeleteVideo   godoc
// @Summary      Delete an existing video
// @Description  Takes a video id and detele it in DB. Return nothing.
// @Tags         videos
// @Produce      json
// @Param        uint   path      uint        true  "VideoId"
// @Success      200    {object}  models.Video
// @Router       /auth/videos/{id} [delete]
// @Security     JWT
func (controller *videoController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 0, 0)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
	}
	_, err = controller.service.Delete(id)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
	}
	ctx.Writer.WriteHeader(http.StatusAccepted)
}

// GetVideo      godoc
// @Summary      Return all videos
// @Description  Takes nothing. Return list of videos.
// @Tags         videos
// @Produce      json
// @Success      200    {array}  models.Video
// @Router       /auth/videos [get]
// @Security     JWT
func (controller *videoController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.service.FindAll())
}

// DeleteVideo   godoc
// @Summary      Find an existing video
// @Description  Takes a video id and return it from DB.
// @Tags         videos
// @Produce      json
// @Param        uint   path      uint        true  "VideoId"
// @Success      200    {object}  models.Video
// @Router       /auth/videos/{id} [get]
// @Security     JWT
func (controller *videoController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not bind or validate data!",
		})
	}
	log.Printf("id is %+v\n", id)
	video := controller.service.FindById(id)
	if video == (models.Video{}) {
		ctx.Writer.WriteHeader(404)
	}
	ctx.JSON(http.StatusOK, video)
}
