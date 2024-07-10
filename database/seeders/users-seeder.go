package seeders

import (
	"fmt"

	"github.com/akposieyefa/open-ai/core/models"
	"github.com/akposieyefa/open-ai/internals"
	"github.com/akposieyefa/open-ai/pkg/hash"
	"github.com/bxcodec/faker/v3"
)

func UserSeedeers() {

	password, err := hash.HashPassword("password")

	if err != nil {
		fmt.Println("Sorry there was an error hashing the password")
	}

	users := []models.User{
		{
			Name:            "Orutu Akposieyefa Williams",
			Email:           "orutu1@gmail.com",
			PhoneNumber:     "+12381000788859",
			IsEmailVerified: true,
			Password:        string(password),
			RoleId:          1,
		},
		{
			Name:            faker.Name(),
			Email:           faker.Email(),
			PhoneNumber:     faker.E164PhoneNumber(),
			IsEmailVerified: true,
			Password:        string(password),
			RoleId:          3,
		},
		{
			Name:            faker.Name(),
			Email:           faker.Email(),
			PhoneNumber:     faker.E164PhoneNumber(),
			IsEmailVerified: true,
			Password:        string(password),
			RoleId:          2,
		},
	}
	internals.DB.Create(&users)

}
