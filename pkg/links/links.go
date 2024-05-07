package links

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aliamerj/dub-go/internal/handler"
	"github.com/aliamerj/dub-go/internal/utils"
)

type LinksService struct {
	Token       string
	WorkspaceId string
	HttpClient  *http.Client
}

func (ls *LinksService) Create(options Options) (*responseOptions, error) {
	if len(options.URL) == 0 {
		return nil, fmt.Errorf("URL is required")
	}
	apiUrl := "https://api.dub.co/links?workspaceId=" + ls.WorkspaceId
	jsonData, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ls.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response responseOptions
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil
}

func (ls *LinksService) Get(options GetOptions) (*responseOptions, error) {
	baseURL := "https://api.dub.co/links/info"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", ls.WorkspaceId)

	// Check and append the provided parameters
	if len(options.Domain) > 0 && len(options.Key) > 0 {
		query.Set("domain", options.Domain)
		query.Set("key", options.Key)
	} else if len(options.LinkId) > 0 {
		query.Set("linkId", options.LinkId)
	} else if len(options.ExternalId) > 0 {
		query.Set("externalId", options.ExternalId)
	} else {
		return nil, fmt.Errorf("either domain and key, linkId or externalId must be provided")
	}
	url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+ls.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response responseOptions
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil

}

func (ls *LinksService) Update(linkId string, options Options) (*responseOptions, error) {
	if len(linkId) == 0 {
		return nil, fmt.Errorf("LinkId should not be  empty")
	}
	apiUrl := "https://api.dub.co/links/" + linkId + "?workspaceId=" + ls.WorkspaceId
	jsonData, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ls.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response responseOptions
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil
}

func (ls *LinksService) Delete(linkId string) (*deleteResponse, error) {
	if len(linkId) == 0 {
		return nil, fmt.Errorf("LinkId should not be  empty")
	}
	apiUrl := "https://api.dub.co/links/" + linkId + "?workspaceId=" + ls.WorkspaceId

	req, err := http.NewRequest("DELETE", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ls.Token)
	res, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response deleteResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil
}
func (ls *LinksService) List(options *GetListOptions) ([]responseOptions, error) {
	baseURL := "https://api.dub.co/links"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", ls.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}

	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+ls.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []responseOptions
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (ls *LinksService) Count(options *GetCountOptions) (int, error) {
	baseURL := "https://api.dub.co/links/count"
	url, err := url.Parse(baseURL)
	if err != nil {
		return 0, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", ls.WorkspaceId)
	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("error fetch count request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+ls.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := ls.HttpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, handler.APIError(resp)
	}
	var response int
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (ls *LinksService) CreateMany(options []Options) ([]responseOptions, error) {
	if len(options) == 0 {
		return nil, fmt.Errorf("No links created – you must provide at least one link.")
	}
	for _, option := range options {
		if len(option.URL) == 0 {
			return nil, fmt.Errorf("URL is required")
		}

	}
	apiUrl := "https://api.dub.co/links/bulk" + "?workspaceId=" + ls.WorkspaceId

	jsonData, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ls.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response []responseOptions
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}
