package model

import (
	"anti-hangmango-web-api/api"
	"log"
	"time"
)

type User struct {
	LoginName string
	Password  string
	AuthToken string
	ExpiredAt int64
}

func NewUser(loginName string, password string) (*User, error) {
	user := new(User)
	user.LoginName = loginName
	user.Password = password
	if err := user.SignUp(); err != nil {
		return nil, err
	}
	if err := user.SignIn(); err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) SignUp() error {
	res, err := api.UserSignUp(user.LoginName, user.Password)
	if err != nil {
		return api.BaseAPIRespErrorHandle(res, err)
	}
	defer res.Body.Close()
	log.Printf("User: %s sign up success\n", user.LoginName)
	return err
}

func (user *User) SignIn() error {
	res, err := api.UserSignIn(user.LoginName, user.Password)
	if err != nil {
		return api.BaseAPIRespErrorHandle(res, err)
	}
	defer res.Body.Close()
	resBodyMap, _ := res.ParseBodyToMap()
	user.AuthToken = resBodyMap["token"].(string)
	user.ExpiredAt = int64(resBodyMap["expired_at"].(float64))
	log.Printf("User: %s sign in success\n", user.LoginName)
	return err
}

func (user *User) IsUserAuth() bool {
	if user.AuthToken == "" || user.ExpiredAt == 0 {
		return false
	}
	if user.ExpiredAt <= time.Now().Unix() {
		return false
	}
	return true
}

func BestUsers(page int64, pageSize int64) error {
	res, err := api.GetBestUsers(page, pageSize)
	resBodyJSON, err := res.ParseBodyToJSON()
	defer res.Body.Close()
	log.Printf("Request Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJSON))
	return err
}
