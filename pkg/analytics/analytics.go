package analytics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aliamerj/dub-go/internal/handler"
	"github.com/aliamerj/dub-go/internal/utils"
)

type AnalyticsService struct {
	Token       string
	WorkspaceId string
	HttpClient  *http.Client
}

func (as *AnalyticsService) Clicks(options *Options) (int, error) {
	baseURL := "https://api.dub.co/analytics/clicks"
	url, err := url.Parse(baseURL)
	if err != nil {
		return 0, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
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

func (as *AnalyticsService) Timeseries(options *Options) ([]timeseriesResponse, error) {
	baseURL := "https://api.dub.co/analytics/timeseries"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []timeseriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) Countries(options *Options) ([]countriesResponse, error) {
	baseURL := "https://api.dub.co/analytics/country"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []countriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) Cities(options *Options) ([]citiesResponse, error) {
	baseURL := "https://api.dub.co/analytics/city"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []citiesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) Devices(options *Options) ([]devicesResponse, error) {
	baseURL := "https://api.dub.co/analytics/device"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []devicesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) Browsers(options *Options) ([]browsersResponse, error) {
	baseURL := "https://api.dub.co/analytics/browser"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []browsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) OS(options *Options) ([]osResponse, error) {
	baseURL := "https://api.dub.co/analytics/os"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []osResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) Referers(options *Options) ([]refererResponse, error) {
	baseURL := "https://api.dub.co/analytics/referer"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []refererResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) TopLinks(options *Options) ([]topLinkResponse, error) {
	baseURL := "https://api.dub.co/analytics/top_links"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []topLinkResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}

func (as *AnalyticsService) TopUrls(options *Options) ([]topUrlsResponse, error) {
	baseURL := "https://api.dub.co/analytics/top_urls"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	query.Set("workspaceId", as.WorkspaceId)

	if options != nil {
		utils.BuildQueryURL(&query, options)
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+as.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := as.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}

	var response []topUrlsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil
}
