package model

import (
	"anti-hangmango-web-api/api"
	"log"
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
	res, err := api.UserSignUp("test1111", "12")
	resBodyJson, err := res.ParseBodyToJSON()
	defer res.Body.Close()
	log.Printf("Request Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJson))
	return err
}

func BestUsers(page int64, pageSize int64) error {
	res, err := api.GetBestUsers(page, pageSize)
	resBodyJson, err := res.ParseBodyToJSON()
	defer res.Body.Close()
	log.Printf("Request Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJson))
	return err
}
