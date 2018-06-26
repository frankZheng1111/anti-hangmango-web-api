package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	urlpkg "net/url"
)

type Response http.Response

func (resp *Response) ParseBodyToMap() (map[string]interface{}, error) {
	resBody := map[string]interface{}{} //= make(map[string]interface{}) //2选1
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
	if resp.StatusCode != 200 {
		err = errors.New("StatusNotOK")
	}
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
	if resp.StatusCode != 200 {
		err = errors.New("StatusNotOK")
	}
	return &resp, err
}

func BaseAPIRespErrorHandle(res *Response, err error) error {
	if err.Error() == "StatusNotOK" {
		resBodyJson, _ := res.ParseBodyToJSON()
		log.Printf("Request Error Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJson))
	}
	return err
}
