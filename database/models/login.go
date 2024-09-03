package models

// This struct is used for logining in
type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
