package models

type User struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   string `json:"status"`
}
