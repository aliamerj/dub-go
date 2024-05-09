package domains

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliamerj/dub-go/internal/handler"
)

type DomainService struct {
	Token       string
	WorkspaceId string
	HttpClient  *http.Client
}

func (ds *DomainService) Add(options Options) (*domainsResponse, error) {
	url := "https://api.dub.co/domains?workspaceId=" + ds.WorkspaceId
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
	req.Header.Set("Authorization", "Bearer "+ds.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ds.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response domainsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil
}

func (ds *DomainService) Update(slug string, options Options) (*domainsResponse, error) {
	if len(slug) == 0 {
		return nil, fmt.Errorf("slug is required")
	}
	url := "https://api.dub.co/domains/" + slug + "?workspaceId=" + ds.WorkspaceId
	jsonData, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ds.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ds.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response domainsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil
}

func (ds *DomainService) Delete(slug string) (*domainsResponse, error) {
	if len(slug) == 0 {
		return nil, fmt.Errorf("slug is required")
	}
	url := "https://api.dub.co/domains/" + slug + "?workspaceId=" + ds.WorkspaceId

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ds.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ds.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response domainsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil
}

func (ds *DomainService) List() ([]domainsResponse, error) {
	url := "https://api.dub.co/domains?workspaceId=" + ds.WorkspaceId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ds.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ds.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response []domainsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}
	return response, nil

}

func (ds *DomainService) SetPrimary(slug string) (*domainsResponse, error) {
	if len(slug) == 0 {
		return nil, fmt.Errorf("slug is required")
	}
	url := "https://api.dub.co/domains/" + slug + "/primary?workspaceId=" + ds.WorkspaceId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ds.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ds.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response domainsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}
	return &response, nil
}

func (ds *DomainService) Transfer(slug string, newWorkspaceId string) (*domainsResponse, error) {
	if len(slug) == 0 {
		return nil, fmt.Errorf("slug is required")
	}
	if len(newWorkspaceId) == 0 {
		return nil, fmt.Errorf("newWorkspaceId is required")
	}
	url := "https://api.dub.co/domains/" + slug + "/transfer?workspaceId=" + ds.WorkspaceId

	body := struct {
		NewWorkspaceId string `json:"newWorkspaceId"`
	}{
		NewWorkspaceId: newWorkspaceId,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ds.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ds.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response domainsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}
	return &response, nil

}
