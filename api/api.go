package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
)

type Response http.Response

func (resp *Response) ParseBodyToMap() (map[string]interface{}, error) {
	var resBody map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&resBody); err != nil {
		return nil, err
	}
	return resBody, nil
}

func (resp *Response) ParseBodyToJSON() ([]byte, error) {
	buf, err := ioutil.ReadAll(resp.Body)
	return buf, err
}

func Post(url string, body map[string]interface{}) (*Response, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	result, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	resp := Response(*result)
	return &resp, err
}

func Get(url string, query map[string]string) (*Response, error) {
	parsedUrl, err := urlpkg.Parse(url)
	if err != nil {
		return nil, err
	}
	ParsedQuery := parsedUrl.Query()
	for key, value := range query {
		ParsedQuery.Add(key, value)
	}
	parsedUrl.RawQuery = ParsedQuery.Encode()
	result, err := http.Get(parsedUrl.String())
	if err != nil {
		return nil, err
	}
	resp := Response(*result)
	return &resp, err
}
