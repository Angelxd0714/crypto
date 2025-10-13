package dto

import "gin-quickstart/internal/utils"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

var validateLogin = utils.ValidateStruct(&LoginRequest{})

func (r *LoginRequest) Validate() error {
	return validateLogin
}
