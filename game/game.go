package game

import (
	"anti-hangmango-web-api/config"
	"anti-hangmango-web-api/model"
	"github.com/satori/go.uuid"
	"log"
	"sync"
)

func Start() {
	wg := sync.WaitGroup{}
	userCount := config.Config.UserCount
	wg.Add(userCount)
	for i := 0; i < userCount; i++ {
		go func() {
			uV4 := uuid.NewV4()
			OnePlayerBegin(uV4.String())
			wg.Done()
		}()
	}
	wg.Wait()
}

func OnePlayerBegin(playerName string) {
	user, err := model.NewUser(playerName, playerName)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("success, user: ", user)
	hangmanCount := config.Config.HangmanCount
	wg := sync.WaitGroup{}
	wg.Add(hangmanCount)
	for i := 0; i < hangmanCount; i++ {
		go func() {
			PlayOneHangman(user)
			wg.Done()
		}()
	}
	wg.Wait()
}

func PlayOneHangman(user *model.User) {
	if hangman, err := model.UserNewHangman(user); err == nil {
		for !hangman.IsFinish() {
			if err = hangman.GuessNextLetter(user); err != nil {
				log.Println(err)
				break
			}
		}
	} else {
		log.Println(err)
	}
}
