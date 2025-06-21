package domain

import "github.com/google/uuid"

type UserRegisterRequest struct {
	UserId   uuid.UUID
	Email    string `json:"email" validate:"required,customEmail"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserRegisterResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
