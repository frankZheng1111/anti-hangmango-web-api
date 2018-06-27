package api

import (
	"anti-hangmango-web-api/config"
	"fmt"
)

func NewHangman(token string) (*Response, error) {
	res, err := Post(config.Config.APIUrl+"/v1/hangmen", map[string]string{"hangmango-auth-token": token}, nil)
	return res, err
}

func HangmanGuessALetter(token string, hangmanId int64, letter string) (*Response, error) {
	url := fmt.Sprintf("%s/v1/hangmen/%d/guess", config.Config.APIUrl, hangmanId)
	header := map[string]string{"hangmango-auth-token": token}
	body := map[string]interface{}{"letter": letter}
	res, err := Post(url, header, body)
	return res, err
}
