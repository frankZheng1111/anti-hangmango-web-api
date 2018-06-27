package model

import (
	"anti-hangmango-web-api/api"
	"log"
)

type Hangman struct {
	Id            int64
	Hp            int8
	Word          string
	StaticLetters []map[rune]int
}

func UserNewHangman(user *User) (*Hangman, error) {
	res, err := api.NewHangman(user.AuthToken)
	if err != nil {
		return nil, api.BaseAPIRespErrorHandle(res, err)
	}
	defer res.Body.Close()
	resBodyMap, _ := res.ParseBodyToMap()
	hangman := new(Hangman)
	hangman.Id = int64(resBodyMap["id"].(float64))
	hangman.Hp = int8(resBodyMap["hp"].(float64))
	hangman.Word = resBodyMap["word"].(string)
	log.Println("New Hangman success: ", hangman)
	return hangman, nil
}
