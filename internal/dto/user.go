package dto

import (
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
)

type UserProfileResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	CityID string `json:"city_id"`
}

type UpdatePasswordRequest struct {
	OldPassword		string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
