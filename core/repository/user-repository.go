package repository

import (
	"fmt"

	"github.com/akposieyefa/open-ai/core/models"
	"github.com/akposieyefa/open-ai/core/types"
	"github.com/akposieyefa/open-ai/internals"
	"github.com/akposieyefa/open-ai/pkg/hash"
)

func CreateUserRepository(name string, email string, phone string, password string, role uint) (*types.UserResponse, error) {

	hashPassword, err := hash.HashPassword(password)

	if err != nil {
		fmt.Println("Sorry there was an error hashing the password")
	}

	user := &models.User{
		Name:        name,
		Email:       email,
		PhoneNumber: phone,
		Password:    string(hashPassword),
		RoleId:      role,
	}

	result := internals.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	userResponse := types.UserResponse{
		ID:              user.ID.String(),
		Name:            user.Name,
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		PhoneNumber:     user.PhoneNumber,
		RoleId:          user.RoleId,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt,
	}

	return &userResponse, nil

}
