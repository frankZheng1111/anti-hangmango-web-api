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

var client *http.Client = &http.Client{}

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

func Post(url string, header map[string]string, body map[string]interface{}) (*Response, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	for key, value := range header {
		req.Header.Add(key, value)
	}
	result, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	resp := Response(*result)
	if resp.StatusCode != 200 {
		err = errors.New("StatusNotOK")
	}
	return &resp, err
}

func Get(url string, header map[string]string, query map[string]string) (*Response, error) {
	parsedUrl, err := urlpkg.Parse(url)
	if err != nil {
		return nil, err
	}
	ParsedQuery := parsedUrl.Query()
	for key, value := range query {
		ParsedQuery.Add(key, value)
	}
	parsedUrl.RawQuery = ParsedQuery.Encode()
	req, err := http.NewRequest("GET", parsedUrl.String(), nil)
	for key, value := range header {
		req.Header.Add(key, value)
	}
	result, err := client.Do(req)
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
