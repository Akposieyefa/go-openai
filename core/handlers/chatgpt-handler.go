package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akposieyefa/open-ai/core/models"
	"github.com/akposieyefa/open-ai/core/services"
	response "github.com/akposieyefa/open-ai/pkg/http"
)

func ChatCompletionHandler(w http.ResponseWriter, r *http.Request) {

	chat := &models.Chat{}

	json.NewDecoder(r.Body).Decode(chat)

	ai, err := services.ChatCompletion(chat.Content)

	if err != nil {
		response.WriteErrorJson("Chat Completion error", err.Error(), false, w, http.StatusBadRequest)
		return
	}

	chatResponse := map[string]string{
		"responseContent": ai.Content,
	}
	response.WriteSuccessJson("Chat retrived from chatgpt succesfully", chatResponse, true, w, http.StatusOK)
}
