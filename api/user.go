package api

import (
	"anti-hangmango-web-api/config"
)

func UserSignUp(loginName string, password string) (*Response, error) {
	res, err := Post(config.Config.ApiUrl+"/v1/users", map[string]interface{}{
		"login_name": loginName,
		"password":   password,
	})
	return res, err
}
