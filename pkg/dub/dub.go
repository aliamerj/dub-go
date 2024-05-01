package dub

import (
	"net/http"
	"time"

	"github.com/aliamerj/dub-go/pkg/links"
)

type Dub struct {
	Links *links.LinksService
}

func NewConfig(token string, workspaceId string) *Dub {
	return &Dub{
		Links: &links.LinksService{
			Token:       token,
			WorkspaceId: workspaceId,
			HttpClient: &http.Client{
				Timeout: time.Second * 30,
			}},
	}
}
