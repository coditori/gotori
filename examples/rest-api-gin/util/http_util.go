package util

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var AccessToken interface{}

func HttpCall(method string, url string, values url.Values, payload io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
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
