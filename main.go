package main

import (
	"anti-hangmango-web-api/model"
	"log"
)

func main() {
	user, err := model.NewUser("test1111", "1111")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("success, user: ", user)
}
