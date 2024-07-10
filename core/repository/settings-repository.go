package repository

import (
	"fmt"
	"net/http"

	"github.com/akposieyefa/open-ai/core/repository/auth"
	"github.com/akposieyefa/open-ai/pkg/hash"
)

func ChangeUserPassword(oldPassword string, password string, r *http.Request) {

	_, err := hash.HashPassword(password)

	_, err := auth.AuthUser(r)
	if err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Println("Sorry there was an error hashing the password")
	}

}
