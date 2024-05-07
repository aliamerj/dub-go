package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		DocURL  string `json:"doc_url"`
	} `json:"error"`
}

func APIError(resp *http.Response) error {
	var apiError apiError
	if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
		return fmt.Errorf("failed to parse API error response: %v", err)
	}
	return fmt.Errorf("Error: %s", apiError.Error.Message)
}
