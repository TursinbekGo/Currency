package models

type Register struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterResponse struct {
	OK bool `json:"ok"`
}
