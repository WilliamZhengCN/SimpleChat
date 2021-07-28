package model

type User struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
