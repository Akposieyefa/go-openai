package handlers

import (
	"net/http"

	"github.com/akposieyefa/open-ai/core/repository"
	response "github.com/akposieyefa/open-ai/pkg/http"
)

// get all banks
func GetAllBanks(w http.ResponseWriter, r *http.Request) {

	banks, err := repository.GetBanks()
	if err != nil {
		response.WriteErrorJson("Sorry unable to get banks", err.Error(), false, w, http.StatusBadRequest)
		return
	}

	response.WriteSuccessJson("Banks fetched successfully", banks, true, w, http.StatusOK)
}
