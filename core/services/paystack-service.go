package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/akposieyefa/open-ai/config"
)

var apiKey = config.Envs.PAYSTACK_SECRET_KEY
var apiUrl = config.Envs.PAYSTACK_URL
var callBackUrl = config.Envs.PAYSTACK_CALLBACK_URL

// initialize payment
func InitializeTransaction(email string, amount int, reference string) (string, error) {

	params := map[string]interface{}{
		"email":        email,
		"amount":       amount,
		"reference":    reference,
		"callback_url": fmt.Sprintf("%s=%s", callBackUrl, reference),
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiUrl+"/transaction/initialize", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("unexpected status code: " + resp.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	authURL, ok := data["data"].(map[string]interface{})["authorization_url"].(string)
	if !ok {
		return "", errors.New("authorization URL not found in response")
	}

	return authURL, nil
}

// verify transaction
func VerifyTransaction(reference string) (bool, error) {

	req, err := http.NewRequest("GET", apiUrl+"/transaction/verify/"+reference, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return false, err
	}

	status, ok := data["data"].(map[string]interface{})["status"].(bool)
	if !ok {
		return false, errors.New("authorization URL not found in response")
	}
	return status, nil
}
