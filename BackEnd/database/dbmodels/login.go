package dbmodels

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequestClient struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterRequestOwner struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Career   string `json:"career"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
