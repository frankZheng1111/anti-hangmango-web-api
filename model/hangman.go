package model

import (
	"anti-hangmango-web-api/api"
	"anti-hangmango-web-api/config"
	"log"
	"strings"
)

type Hangman struct {
	Id            int64
	Hp            int8
	Word          string
	StaticLetters []map[rune]int
	LettersCount  map[rune]int
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
	log.Printf("New Hangman success: Id: %d, Word: %s\n", hangman.Id, hangman.Word)
	return hangman, nil
}

func (hangman *Hangman) InitDictionary() {
	const POSSIBLE_LETTER_COUNT = 36
	wordLen := len([]rune(hangman.Word))
	hangman.StaticLetters = make([]map[rune]int, wordLen)
	hangman.LettersCount = make(map[rune]int, POSSIBLE_LETTER_COUNT)
	for i := 0; i < wordLen; i++ {
		hangman.StaticLetters[i] = make(map[rune]int, POSSIBLE_LETTER_COUNT)
	}
	for _, word := range config.Config.Hangman.Dictionary {
		letterRunes := []rune(word)
		if wordLen != len(letterRunes) {
			continue
		}
		for index, letterRune := range letterRunes {
			hangman.StaticLetters[index][letterRune]++
			hangman.LettersCount[letterRune]++
		}
	}
}

func (hangman *Hangman) UpdateRemainLetter(letter string) {
	letterRune := []rune(letter)[0]
	correctPositions := []int{}
	for index, wordLetter := range hangman.Word {
		if letterRune == wordLetter {
			correctPositions = append(correctPositions, index)
		}
	}
	for _, correctPosition := range correctPositions {
		for colLetterRune, count := range hangman.StaticLetters[correctPosition] {
			hangman.LettersCount[colLetterRune] -= count
		}
	}
	hangman.LettersCount[letterRune] = 0
}

func (hangman *Hangman) MostInStaticLetters() string {
	var mostLetter rune
	for letter, count := range hangman.LettersCount {
		if mostLetter == 0 {
			mostLetter = letter
		} else if count > hangman.LettersCount[mostLetter] {
			mostLetter = letter
		}
	}
	return string(mostLetter)
}

func (hangman *Hangman) GuessNextLetter(user *User) error {
	guessLetter := hangman.MostInStaticLetters()
	res, err := api.HangmanGuessALetter(user.AuthToken, hangman.Id, hangman.MostInStaticLetters())
	if err != nil {
		return api.BaseAPIRespErrorHandle(res, err)
	}
	defer res.Body.Close()
	resBodyMap, _ := res.ParseBodyToMap()
	hangman.Hp = int8(resBodyMap["hp"].(float64))
	hangman.Word = resBodyMap["word"].(string)
	hangman.UpdateRemainLetter(guessLetter)
	log.Printf("Hangman Guess Letter Id: %d, Word: %s, Letter: %s, Hp: %d\n", hangman.Id, hangman.Word, guessLetter, hangman.Hp)
	return nil
}

func (hangman *Hangman) IsWin() bool {
	return !strings.Contains(hangman.Word, "*")
}

func (hangman *Hangman) IsFinish() bool {
	return hangman.IsWin() || hangman.Hp == 0
}
