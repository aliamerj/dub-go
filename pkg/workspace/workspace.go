package workspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliamerj/dub-go/internal/handler"
)

type WorkspaceServices struct {
	Token      string
	HttpClient *http.Client
}

func (ws *WorkspaceServices) Create(options Options) (*responseWorkspace, error) {
	url := "https://api.dub.co/workspaces"
	if len(options.Name) == 0 {
		return nil, fmt.Errorf("Name is required")
	}
	if len(options.Slug) == 0 {
		return nil, fmt.Errorf("Slug is required")
	}

	jsonData, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ws.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ws.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response responseWorkspace
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil

}

func (ws *WorkspaceServices) Get(idOrSlug string) (*responseWorkspace, error) {
	if len(idOrSlug) == 0 {
		return nil, fmt.Errorf("idOrSlug is required")
	}
	url := "https://api.dub.co/workspaces/" + idOrSlug
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ws.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ws.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response responseWorkspace
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}
	return &response, nil

}

func (ws *WorkspaceServices) List() ([]responseWorkspace, error) {

	url := "https://api.dub.co/workspaces"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ws.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ws.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response []responseWorkspace
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}
	return response, nil

}
