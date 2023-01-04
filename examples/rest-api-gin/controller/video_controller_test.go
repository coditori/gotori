package controller

import (
	"net/http"
	"net/url"
	"rest/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenCorrectRequest_WhenUserLoggedIn_ThenCreateVideoWillSucceed(t *testing.T) {
	token := util.GetToken()
	payload := strings.NewReader(`{
		"title": "Develop",
		"description": "Java, Pyhton, JS",
		"url": "https://ghitub.com/coditori",
		"author": {
			"firstName": "Ario",
			"lastName": "Afrashteh",
			"age": 404,
			"email": "sample@gmail.com"
		}
	}`)

	res, err := util.HttpCall("POST", "http://localhost:8000/auth/videos", url.Values{}, payload, token)
	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, http.StatusCreated)
}

func TestGivenDuplicateRequest_WhenUrlIsDuplicateFoAnotherUser_ThenCreateVideoWillFail(t *testing.T) {
	token := util.GetToken()
	payload := strings.NewReader(`{
		"title": "Develop",
		"description": "Java, Pyhton, JS",
		"url": "https://ghitub.com/coditori",
		"author": {
			"firstName": "Ario",
			"lastName": "Afrashteh",
			"age": 404,
			"email": "sample@gmail.com"
		}
	}`)

	res, err := util.HttpCall("POST", util.BaseUrl+"/auth/videos", url.Values{}, payload, token)
	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestGivenNothing_WhenTryingToGetAllVideos_ThenOneVideoShouldReturnSuccessfully(t *testing.T) {
	token := util.GetToken()
	res, err1 := util.HttpCall("GET", util.BaseUrl+"/auth/videos", url.Values{}, nil, token)
	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	mappedResponse, err2 := util.ParseHttpResponseToArray(res)
	assert.Equal(t, err2, nil)
	// println(mappedResponse)
	assert.True(t, len(mappedResponse) == 1)
}
