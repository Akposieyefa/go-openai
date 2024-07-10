package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akposieyefa/open-ai/core/repository/auth"
	"github.com/akposieyefa/open-ai/core/types"
	"github.com/akposieyefa/open-ai/pkg"
	response "github.com/akposieyefa/open-ai/pkg/http"
)

// login user
func Login(w http.ResponseWriter, r *http.Request) {

	payload := types.UserLoginPayload{}
	json.NewDecoder(r.Body).Decode(&payload)

	err := pkg.Validate(payload)
	if err != nil {
		response.WriteErrorJson("Sorry json is not valid", err.Error(), false, w, http.StatusBadRequest)
		return
	}

	token, err := auth.GetAuthToken(payload.Email, payload.Password)
	if err != nil {
		response.WriteErrorJson(err.Error(), err.Error(), false, w, http.StatusBadRequest)
		return
	}

	response.WriteSuccessJson("User logged in successfully", token, true, w, http.StatusCreated)

}

// get logged in users profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := auth.AuthUser(r)
	if err != nil {
		response.WriteErrorJson(err.Error(), err.Error(), false, w, http.StatusBadRequest)
		return
	}

	response.WriteSuccessJson("User profile pulled successfully", profile, true, w, http.StatusCreated)
}
