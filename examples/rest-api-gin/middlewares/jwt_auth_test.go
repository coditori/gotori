package middlewares

import (
	"net/http"
	"net/url"
	"rest/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenCorrectRequest_WhenUserIsNotLoggedIn_ThenLoginWillSucceed(t *testing.T) {
	payload := strings.NewReader(`{
		"username": "admin",
		"password": "admin"
	}`)

	res, err1 := util.HttpCall("POST", "http://localhost:8000/login", url.Values{}, payload)

	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	mappedResponse, err2 := util.ParseHttpResponseToMap(res)

	assert.Equal(t, err2, nil)
	assert.NotEmpty(t, mappedResponse)
	util.AccessToken = mappedResponse["token"]
}

func TestGivenWronogRequest_WhenUserIsNotLoggedIn_ThenLoginWillFail(t *testing.T) {
	payload := strings.NewReader(`{
		"username": "wrongUsername",
		"password": ""
	}`)

	res, err1 := util.HttpCall("POST", "http://localhost:8000/login", url.Values{}, payload)

	assert.Equal(t, err1, nil)
	assert.Equal(t, res.StatusCode, http.StatusUnauthorized)
}
