package qr

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/aliamerj/dub-go/internal/handler"
)

type QrService struct {
	Token      string
	HttpClient *http.Client
}

func (qs *QrService) Get(option *Options) ([]byte, error) {
	baseURL := "https://api.dub.co/qr"
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	query := url.Query()
	if option != nil {
		if option.URL != "" {
			query.Add("url", option.URL)
		}
		if option.IncludeMargin {
			query.Add("includeMargin", "true")
		}
		if option.BgColor != "" {
			query.Add("bgColor", option.BgColor)
		}
		if option.FgColor != "" {
			query.Add("fgColor", option.FgColor)
		}
		if option.Level != "" {
			query.Add("level", string(option.Level))
		}
		if option.Size > 0 {
			query.Add("size", fmt.Sprintf("%d", option.Size))
		}
	}
	url.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error fetch QR: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+qs.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := qs.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handler.APIError(resp)
	}
	png, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read qr response: %v", err)
	}

	return png, nil

}
