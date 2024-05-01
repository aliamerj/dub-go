package links

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliamerj/dub-go/internal/handler"
)

type LinksService struct {
	Token       string
	WorkspaceId string
	HttpClient  *http.Client
}

func (ls *LinksService) Create(requestOptions RequestOptions) (*RequestOptions, error) {
	if len(requestOptions.URL) == 0 {
		return nil, fmt.Errorf("URL is required")
	}
	apiUrl := "https://api.dub.co/links?workspaceId=" + ls.WorkspaceId
	jsonData, err := json.Marshal(requestOptions)
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
	if res.StatusCode != 200 {
		return nil, handler.APIError(res)
	}
	var response RequestOptions
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil
}
