package metatags

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliamerj/dub-go/internal/handler"
)

type MetatagsService struct {
	HttpClient *http.Client
}

func (ms *MetatagsService) Get(url string) (*resMetatags, error) {
	if len(url) == 0 {
		return nil, fmt.Errorf("url is required")
	}
	urlAPI := "https://api.dub.co/metatags?url=" + url
	req, err := http.NewRequest("GET", urlAPI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := ms.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, handler.APIError(res)
	}
	var response resMetatags
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode successful response: %v", err)
	}

	return &response, nil

}
