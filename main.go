package main

import (
	"anti-hangmango-web-api/model"
	"log"
)

func main() {
	model.BestUsers(1, 30)
	user := model.NewUser("test1111", "1111")
	err := user.SignUp()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("success")
}
