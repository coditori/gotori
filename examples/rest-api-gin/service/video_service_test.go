package service

import (
	"rest/models"
	"rest/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
)

var (
	video models.Video = models.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
	videoRepository  repository.VideoRepository = repository.New()
	videoServiceImpl VideoService               = New(videoRepository)
)

func givenVideo_whenSaveVideo_thenReturnVideoObject(t *testing.T) {
	video = videoServiceImpl.Save(video)
	videos := videoServiceImpl.FindAll()

	firstVideo := videos[0]
	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)

}

func givenNothing_whenGetAllVideos_thenReturnVideosList(t *testing.T) {

	videoServiceImpl.Save(video)
	videos := videoServiceImpl.FindAll()

	firstVideo := videos[0]
	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
