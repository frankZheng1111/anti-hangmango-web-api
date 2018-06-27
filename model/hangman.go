package model

import (
	"anti-hangmango-web-api/api"
	"anti-hangmango-web-api/config"
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
	hangman.InitDictionary()
	log.Println("New Hangman success: ", hangman)
	return hangman, nil
}

func (hangman *Hangman) InitDictionary() {
	wordLen := len([]rune(hangman.Word))
	hangman.StaticLetters = make([]map[rune]int, wordLen)
	for i := 0; i < wordLen; i++ {
		hangman.StaticLetters[i] = make(map[rune]int, 26)
	}
	for _, word := range config.Config.Hangman.Dictionary {
		letterRunes := []rune(word)
		if wordLen != len(letterRunes) {
			continue
		}
		for index, letterRune := range letterRunes {
			hangman.StaticLetters[index][letterRune]++
		}
	}
}
