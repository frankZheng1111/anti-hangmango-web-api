package model

type User struct {
	LoginInfo LoginInfo
	authToken string
}

type LoginInfo struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}
