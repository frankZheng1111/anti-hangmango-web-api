package api

import (
	"anti-hangmango-web-api/config"
	"strconv"
)

func UserSignUp(loginName string, password string) (*Response, error) {
	res, err := Post(config.Config.ApiUrl+"/v1/users", map[string]interface{}{
		"login_name": loginName,
		"password":   password,
	})
	return res, err
}

func UserSignIn(loginName string, password string) (*Response, error) {
	res, err := Post(config.Config.ApiUrl+"/v1/users/signin", map[string]interface{}{
		"login_name": loginName,
		"password":   password,
	})
	return res, err
}

func GetBestUsers(page int64, pageSize int64) (*Response, error) {
	pageStr := strconv.FormatInt(page, 2)
	pageSizeStr := strconv.FormatInt(pageSize, 10)
	res, err := Get(config.Config.ApiUrl+"/v1/users/best-users", map[string]string{
		"page":     pageStr,
		"pageSize": pageSizeStr,
	})
	return res, err
}
