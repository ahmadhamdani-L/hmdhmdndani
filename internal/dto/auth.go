package dto

import (
	"lion-super-app/internal/model"
	res "lion-super-app/pkg/util/response"
)

// Login
type AuthLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token                      string `json:"token"`
	model.UserEntityModel
}
type AuthLoginResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AuthLoginResponse `json:"data"`
	} `json:"body"`
}

// Register
type AuthRegisterRequest struct {
	model.UserEntityModel
}
type AuthRegisterResponse struct {
	model.UserEntityModel
}
type AuthRegisterResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data AuthRegisterResponse `json:"data"`
	} `json:"body"`
}