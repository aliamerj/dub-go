package analytics

type interval string

const (
	Interval1Hour  interval = "1h"
	Interval24Hour interval = "24h"
	Interval7Day   interval = "7d"
	Interval30Day  interval = "30d"
	Interval90Day  interval = "90d"
	IntervalAll    interval = "all"
)

type Options struct {
	Domain     string   `json:"domain,omitempty"`
	Key        string   `json:"key,omitempty"`
	LinkId     string   `json:"linkId,omitempty"`
	ExternalId string   `json:"externalId,omitempty"`
	Interval   interval `json:"interval,omitempty"`
	Country    string   `json:"country,omitempty"`
	City       string   `json:"city,omitempty"`
	Device     string   `json:"device,omitempty"`
	Browser    string   `json:"browser,omitempty"`
	Os         string   `json:"os,omitempty"`
	Referer    string   `json:"referer,omitempty"`
	Url        string   `json:"url,omitempty"`
	TagId      string   `json:"tagId,omitempty"`
	Qr         bool     `json:"qr,omitempty"`
	Root       bool     `json:"root,omitempty"`
}
