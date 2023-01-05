package controller

import (
	"net/http"
	"net/url"
	"rest/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleVideo = `{
	"title": "Java Development",
	"description": "Java, Pyhton, JS",
	"url": "https://ghitub.com/coditori",
	"author": {
		"firstName": "Ario",
		"lastName": "Afrashteh",
		"age": 99,
		"email": "sample@gmail.com"
	}
}`

func TestGivenCorrectRequest_WhenUserLoggedIn_ThenCreateVideoWillSucceed(t *testing.T) {
	token := util.GetToken()
	payload := strings.NewReader(sampleVideo)

	res, err := util.HttpCall("POST", util.BaseUrl+"/auth/videos", url.Values{}, payload, token)
	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, http.StatusCreated)
}

// test title with one letter word

func TestGivenDuplicateRequest_WhenUrlIsDuplicateFoAnotherUser_ThenCreateVideoWillFail(t *testing.T) {
	token := util.GetToken()
	payload := strings.NewReader(sampleVideo)

	res, err := util.HttpCall("POST", util.BaseUrl+"/auth/videos", url.Values{}, payload, token)
	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

// update

func TestGivenVideoId_WhenTryingToGetOneVideo_ThenOneVideoShouldReturnSuccessfully(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("GET", util.BaseUrl+"/auth/videos/1", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	mappedResponse, err2 := util.ParseHttpResponseToMap(res)
	assert.Equal(t, err2, nil)
	assert.Equal(t, "Java Development", mappedResponse["title"])
}

func TestGivenNothing_WhenTryingToGetAllVideos_ThenOneVideoShouldReturnSuccessfully(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("GET", util.BaseUrl+"/auth/videos", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	mappedResponse, err2 := util.ParseHttpResponseToArray(res)
	assert.Equal(t, err2, nil)
	assert.True(t, len(mappedResponse) == 1)
}

func TestGivenVideoId_WhenTryingToDeleteOneVideo_ThenOneVideoWillDeletedSuccessfully(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("DELETE", util.BaseUrl+"/auth/videos/1", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusAccepted)
}

func TestGivenVideoId_WhenTryingToDeleteNotExistingVideo_ThenDeletionWIllFail(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("DELETE", util.BaseUrl+"/auth/videos/1", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}

func TestGivenStringVideoId_WhenTryingToDeleteOneVideo_ThenDeletionWIllFail(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("DELETE", util.BaseUrl+"/auth/videos/one", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestGivenNothing_WhenTryingToGetEmptyListOfVideos_ThenProcessWIllSucceed(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("GET", util.BaseUrl+"/auth/videos", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	mappedResponse, err2 := util.ParseHttpResponseToArray(res)
	assert.Equal(t, err2, nil)
	assert.True(t, len(mappedResponse) == 0)
}

func TestGivenVideoId_WhenTryingToGetDeletedVideo_ThenProcessWillFail(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("GET", util.BaseUrl+"/auth/videos/1", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}

func TestGivenVideoId_WhenTryingToGetVideoByStringId_ThenProcessWillFail(t *testing.T) {
	token := util.GetToken()

	res, err1 := util.HttpCall("GET", util.BaseUrl+"/auth/videos/ONE", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}
