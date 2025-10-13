package dto

import "gin-quickstart/internal/utils"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (r *RegisterRequest) Validate() error {
	return utils.ValidateStruct(r)
}
