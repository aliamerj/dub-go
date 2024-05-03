package links

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aliamerj/dub-go/internal/handler"
)

type LinksService struct {
	Token       string
	WorkspaceId string
	HttpClient  *http.Client
}

func (ls *LinksService) Create(createOptions RequestOptions) (*responseOptions, error) {
	if len(createOptions.URL) == 0 {
		return nil, fmt.Errorf("URL is required")
	}
	apiUrl := "https://api.dub.co/links?workspaceId=" + ls.WorkspaceId
	jsonData, err := json.Marshal(createOptions)
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

func (ls *LinksService) Get(getOptions GetOptions) (*responseOptions, error) {
	baseURL := "https://api.dub.co/links/info"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", ls.WorkspaceId)

	// Check and append the provided parameters
	if len(getOptions.Domain) > 0 && len(getOptions.Key) > 0 {
		query.Set("domain", getOptions.Domain)
		query.Set("key", getOptions.Key)
	} else if len(getOptions.LinkId) > 0 {
		query.Set("linkId", getOptions.LinkId)
	} else if len(getOptions.ExternalId) > 0 {
		query.Set("externalId", getOptions.ExternalId)
	} else {
		return nil, fmt.Errorf("either domain and key, linkId or externalId must be provided")
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

	var response responseOptions
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil

}

func (ls *LinksService) Update(linkId string, requsetOptions RequestOptions) (*responseOptions, error) {
	if len(linkId) == 0 {
		return nil, fmt.Errorf("LinkId should not be  empty")
	}
	apiUrl := "https://api.dub.co/links/" + linkId + "?workspaceId=" + ls.WorkspaceId
	jsonData, err := json.Marshal(requsetOptions)
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
