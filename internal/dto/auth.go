package dto

import (
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
)

// Login
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Register
type AuthRegisterRequest struct {
	Name    	string `json:"name" validate:"required"`
	Email    	string `json:"email" validate:"required,email"`
	Password 	string `json:"password" validate:"required"`
	Role    	string `json:"role" validate:"required"`
	CityID  	string `json:"city_id" validate:"required"`
}
type AuthRegisterResponse struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}
