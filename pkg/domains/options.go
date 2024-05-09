package domains

type domainType string

const (
	Redirect domainType = "redirect"
	Rewrite  domainType = "rewrite"
)

type Options struct {
	Slug        string     `json:"slug,omitempty"`
	Type        domainType `json:"type,omitempty"`
	Target      *string    `json:"target,omitempty"`
	ExpiredUrl  *string    `json:"expiredUrl,omitempty"`
	Archived    bool       `json:"archived,omitempty"`
	Placeholder *string    `json:"placeholder,omitempty"`
}

type domainsResponse struct {
	ID          string  `json:"id"`
	Slug        string  `json:"slug"`
	Verified    bool    `json:"verified"`
	Primary     bool    `json:"primary"`
	Archived    bool    `json:"archived"`
	Placeholder *string `json:"placeholder"`
	ExpiredUrl  *string `json:"expiredUrl"`
	Target      *string `json:"target"`
	Type        string  `json:"type"`
	Clicks      int     `json:"clicks"`
}

type deleteDomains struct {
	Slug string `json:"slug"`
}
