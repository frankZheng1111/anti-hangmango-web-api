package model

import (
	"anti-hangmango-web-api/api"
	"anti-hangmango-web-api/config"
	"log"
	"strings"
)

const POSSIBLE_LETTER_COUNT = 26

type Hangman struct {
	Id             int64
	Hp             int8
	Word           string
	LettersCount   map[rune]int
	Dictionary     [][]rune
	GuessedLetters map[rune]struct{}
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
	hangman.LettersCount = make(map[rune]int, POSSIBLE_LETTER_COUNT)
	hangman.GuessedLetters = make(map[rune]struct{}, POSSIBLE_LETTER_COUNT)
	hangman.InitDictionary()
	log.Printf("New Hangman success: Id: %d, Word: %s\n", hangman.Id, hangman.Word)
	return hangman, nil
}

func (hangman *Hangman) InitDictionary() {
	wordLen := len([]rune(hangman.Word))
	for _, word := range config.Config.Hangman.Dictionary {
		letterRunes := []rune(word)
		if wordLen != len(letterRunes) {
			continue
		}
		hangman.Dictionary = append(hangman.Dictionary, letterRunes)
		hangman.UpdateLettersCount(letterRunes)
	}
}

func (hangman *Hangman) UpdateLettersCount(letterRunes []rune) {
	for _, letterRune := range letterRunes {
		if _, ok := hangman.GuessedLetters[letterRune]; !ok && string(letterRune) != "-" { // 仅统计未猜过的字母
			hangman.LettersCount[letterRune]++
		}
	}
}

func (hangman *Hangman) UpdateRemainLetter(letter string) {
	hangman.LettersCount = make(map[rune]int, POSSIBLE_LETTER_COUNT)
	letterRune := []rune(letter)[0]
	correctPositions := make(map[int]struct{})
	for index, wordLetter := range hangman.Word {
		if letterRune == wordLetter {
			correctPositions[index] = struct{}{}
		}
	}
	for index := 0; index < len(hangman.Dictionary); index++ {
		wordRunes := hangman.Dictionary[index]
		needAppend := true
		if len(correctPositions) > 0 { // 若该字母猜对了, 留下仅该位置正确的词
			for letterIndex, wordRune := range wordRunes {
				if _, ok := correctPositions[letterIndex]; ok && (wordRune != letterRune) || (!ok && wordRune == letterRune) {
					needAppend = false
					break
				}
			}
		} else if strings.Contains(string(wordRunes), letter) { // 若该字母猜错了, 移除包含该字母的词
			needAppend = false
		}
		if needAppend {
			hangman.UpdateLettersCount(wordRunes)
		} else {
			hangman.RemoveWordInDictionary(index)
			index--
		}
	}
}

func (hangman *Hangman) RemoveWordInDictionary(index int) {
	hangman.Dictionary = append(hangman.Dictionary[0:index], hangman.Dictionary[index+1:]...)
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
	res, err := api.HangmanGuessALetter(user.AuthToken, hangman.Id, guessLetter)
	if err != nil {
		return api.BaseAPIRespErrorHandle(res, err)
	}
	hangman.GuessedLetters[[]rune(guessLetter)[0]] = struct{}{}
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
