package models

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
