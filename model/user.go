package model

import (
	"anti-hangmango-web-api/api"
	"log"
)

type User struct {
	LoginName string
	Password  string
	authToken string
	expiredAt int64
}

func NewUser(loginName string, password string) (*User, error) {
	user := new(User)
	user.LoginName = loginName
	user.Password = password
	if err := user.SignUp(); err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) SignUp() error {
	res, err := api.UserSignUp(user.LoginName, user.Password)
	resBodyJson, err := res.ParseBodyToJSON()
	defer res.Body.Close()
	log.Printf("Request Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJson))
	log.Printf("User: %s sign up success\n", user.LoginName)
	return err
}

func BestUsers(page int64, pageSize int64) error {
	res, err := api.GetBestUsers(page, pageSize)
	resBodyJson, err := res.ParseBodyToJSON()
	defer res.Body.Close()
	log.Printf("Request Response: code: %d, body: %v\n", res.StatusCode, string(resBodyJson))
	return err
}
