package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    uint   `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
