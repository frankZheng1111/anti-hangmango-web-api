package model

import (
	"anti-hangmango-web-api/api"
	"log"
)

func UserNewHangman(user *User) {
	res, _ := api.NewHangman(user.AuthToken)
	resBodyJSON, _ := res.ParseBodyToJSON()
	log.Println("New Hangman success: ", string(resBodyJSON))
	return
}
