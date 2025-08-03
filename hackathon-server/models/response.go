package models

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}