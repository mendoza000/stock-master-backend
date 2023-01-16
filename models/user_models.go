package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	UserRange int    `json:"user_range"`
}
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}