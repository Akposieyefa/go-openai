package repository

import (
	"fmt"
	"net/http"

	"github.com/akposieyefa/open-ai/core/repository/auth"
	"github.com/akposieyefa/open-ai/pkg/hash"
)

func ChangeUserPassword(oldPassword string, password string, r *http.Request) error {

	_, err := hash.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing the password:", err)
		return err
	}

	_, err = auth.AuthUser(r)
	if err != nil {
		fmt.Println("Authentication failed:", err)
		return err
	}

	// Proceed with changing the password in the DB (not implemented here)
	fmt.Println("Password change logic goes here")

	return nil
}
