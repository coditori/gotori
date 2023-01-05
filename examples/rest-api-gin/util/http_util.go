package util

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var AccessToken string
var BaseUrl = "http://localhost:8000"

func HttpCall(method string, url string, values url.Values, payload io.Reader, token string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}

func ParseHttpResponseToMap(res *http.Response) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	x := map[string]interface{}{}
	json.Unmarshal([]byte(body), &x)
	return x, nil
}

func ParseHttpResponseToArray(res *http.Response) ([]map[string]interface{}, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	x := []map[string]interface{}{}
	json.Unmarshal([]byte(body), &x)
	return x, nil
}

func GetToken() string {
	if AccessToken == "" {
		Login()
	}
	return AccessToken
}

func Login() {
	payload := strings.NewReader(`{
			"username": "admin",
			"password": "admin"
		}`)

	res, httpCallErr := HttpCall("POST", "http://localhost:8000/login", url.Values{}, payload, "")
	if httpCallErr != nil {
		log.Printf("Could not log in %+v\n", httpCallErr)
	}
	mappedResponse, parseResponseErr := ParseHttpResponseToMap(res)
	if parseResponseErr != nil {
		log.Printf("Could not parse login response %+v\n", httpCallErr)
	}
	AccessToken = mappedResponse["token"].(string)
}
