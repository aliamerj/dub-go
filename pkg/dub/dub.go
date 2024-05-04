package dub

import (
	"net/http"
	"time"

	"github.com/aliamerj/dub-go/pkg/links"
	"github.com/aliamerj/dub-go/pkg/qr"
)

type Dub struct {
	Links *links.LinksService
	QR    *qr.QrService
}

func NewConfig(token string, workspaceId string) *Dub {
	return &Dub{
		Links: &links.LinksService{
			Token:       token,
			WorkspaceId: workspaceId,
			HttpClient: &http.Client{
				Timeout: time.Second * 30,
			}},
		QR: &qr.QrService{
			Token: token,
			HttpClient: &http.Client{
				Timeout: time.Second * 30,
			},
		},
	}
}
