package seeders

import (
	"github.com/akposieyefa/open-ai/core/models"
	"github.com/akposieyefa/open-ai/internals"
)

func RoleSeeders() {

	roles := []models.Role{
		{
			Name:        "Admin",
			Description: "This is the admin role",
		},
		{
			Name:        "Customer",
			Description: "This is the customer role",
		},
		{
			Name:        "Developer",
			Description: "This is the developer role",
		},
	}
	internals.DB.Create(&roles)

}
