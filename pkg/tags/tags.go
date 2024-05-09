package tags

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliamerj/dub-go/internal/handler"
)

type TagsServices struct {
	Token       string
	WorkspaceId string
	HttpClient  *http.Client
}

func (ts *TagsServices) Create(options Options) (*resTag, error) {
	url := "https://api.dub.co/tags?workspaceId=" + ts.WorkspaceId

	if len(options.Tag) == 0 {
		return nil, fmt.Errorf("tag is required")
	}

	jsonData, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ts.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ts.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response resTag
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil

}

func (ts *TagsServices) List() ([]resTag, error) {
	url := "https://api.dub.co/tags?workspaceId=" + ts.WorkspaceId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ts.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := ts.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response []resTag
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return response, nil

}
