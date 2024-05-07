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
