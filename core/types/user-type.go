package types

import (
	"time"

	"gorm.io/gorm"
)

type UserRegistrationPayload struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Role        uint   `json:"role" validate:"required"`
}

type UserUpdatePayload struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type UserPasswordChangePayload struct {
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	PhoneNumber     string         `json:"phone"`
	IsEmailVerified bool           `json:"is_email_verified"`
	RoleId          uint           `json:"role"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
