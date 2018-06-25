package model

import (
	"anti-hangmango-web-api/api"
	"anti-hangmango-web-api/config"
	"log"
	"strconv"
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
	res, err := api.Post(config.Config.ApiUrl+"/v1/users", map[string]interface{}{
		"login_name": "test1111",
		"password":   "test1111",
	})
	resBodyJson, err := res.ParseBodyToJSON()
	defer res.Body.Close()
	log.Printf("Request Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJson))
	return err
}

func BestUsers(page int64, pageSize int64) error {
	pageStr := strconv.FormatInt(page, 2)
	pageSizeStr := strconv.FormatInt(pageSize, 10)
	res, err := api.Get(config.Config.ApiUrl+"/v1/users/best-users", map[string]string{
		"page":     pageStr,
		"pageSize": pageSizeStr,
	})
	resBodyJson, err := res.ParseBodyToJSON()
	defer res.Body.Close()
	log.Printf("Request Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJson))
	return err
}
