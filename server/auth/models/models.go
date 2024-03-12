package models

type EmailLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsernameLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
