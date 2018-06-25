package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	StatusCode int
	Body       map[string]interface{}
}

func Post(url string, body map[string]interface{}) (*ApiResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	result, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(reqBody))
	defer result.Body.Close()
	if err != nil {
		return nil, err
	}
	var resBody map[string]interface{}
	decoder := json.NewDecoder(result.Body)
	if err = decoder.Decode(&resBody); err != nil {
		return nil, err
	}
	return &ApiResponse{result.StatusCode, resBody}, err
}

// func Get(url string, query map[string]interface{}) (map[string]interface{}, error) {
// }
