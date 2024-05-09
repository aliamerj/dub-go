package dub

import (
	"net/http"
	"time"

	"github.com/aliamerj/dub-go/pkg/analytics"
	"github.com/aliamerj/dub-go/pkg/links"
	"github.com/aliamerj/dub-go/pkg/qr"
	"github.com/aliamerj/dub-go/pkg/workspace"
)

type Dub struct {
	Links     *links.LinksService
	QR        *qr.QrService
	Analytics *analytics.AnalyticsService
	Workspace *workspace.WorkspaceServices
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
		Analytics: &analytics.AnalyticsService{
			Token:       token,
			WorkspaceId: workspaceId,
			HttpClient: &http.Client{
				Timeout: time.Second * 30,
			},
		},
		Workspace: &workspace.WorkspaceServices{
			Token: token,
			HttpClient: &http.Client{
				Timeout: time.Second * 30,
			},
		},
	}
}
