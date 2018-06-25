package model

import (
	"anti-hangmango-web-api/config"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	LoginInfo LoginInfo
	authToken string
}

type LoginInfo struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}

func NewUser(loginName string, password string) *User {
	user := new(User)
	user.LoginInfo.LoginName = loginName
	user.LoginInfo.Password = password
	return user
}

func (user *User) SignUp() error {
	reqBody, err := json.Marshal(user.LoginInfo)
	if err != nil {
		return err
	}
	result, err := http.Post(config.Config.ApiUrl+"/v1/users", "application/json;charset=utf-8", bytes.NewBuffer(reqBody))
	defer result.Body.Close()
	log.Println("Request Post body", string(reqBody))
	body, err := ioutil.ReadAll(result.Body)
	log.Printf("Request Response: code: %d, body: %v\n", result.StatusCode, string(body))
	return err
}
