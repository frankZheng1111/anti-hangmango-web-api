package main

import (
	"anti-hangmango-web-api/model"
	"log"
)

func main() {
	user := model.NewUser("test1111", "1111")
	err := user.SignUp()
	log.Fatalln(err)
	log.Println("success")
}
