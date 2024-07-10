package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akposieyefa/open-ai/core/repository"
	"github.com/akposieyefa/open-ai/core/types"
	"github.com/akposieyefa/open-ai/pkg"
	response "github.com/akposieyefa/open-ai/pkg/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	payload := types.UserRegistrationPayload{}
	json.NewDecoder(r.Body).Decode(&payload)

	err := pkg.Validate(payload)
	if err != nil {
		response.WriteErrorJson("Sorry json is not valid", err.Error(), false, w, http.StatusBadRequest)
		return
	}

	user, err := repository.CreateUserRepository(payload.Name, payload.Email, payload.PhoneNumber, payload.Password, payload.Role)
	if err != nil {
		response.WriteErrorJson("Sorry unable to create user account", err.Error(), false, w, http.StatusBadRequest)
		return
	}

	response.WriteSuccessJson("Account created successfully", user, true, w, http.StatusCreated)
}
