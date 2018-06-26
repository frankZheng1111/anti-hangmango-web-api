package api

import (
	"anti-hangmango-web-api/config"
)

func NewHangman(token string) (*Response, error) {
	res, err := Post(config.Config.APIUrl+"/v1/hangmen", map[string]string{"hangmango-auth-token": token}, nil)
	return res, err
}
