package links

type Options struct {
	URL         string            `json:"url,omitempty"`
	Domain      string            `json:"domain,omitempty"`
	Key         string            `json:"key,omitempty"`
	ExternalID  string            `json:"externalId,omitempty"`
	Prefix      string            `json:"prefix,omitempty"`
	Archived    bool              `json:"archived,omitempty"`
	PublicStats bool              `json:"publicStats,omitempty"`
	TagID       string            `json:"tagId,omitempty"`
	TagIDs      []string          `json:"tagIds,omitempty"`
	TagNames    string            `json:"tagNames,omitempty"`
	Comments    string            `json:"comments,omitempty"`
	ExpiresAt   string            `json:"expiresAt,omitempty"`
	ExpiredURL  string            `json:"expiredUrl,omitempty"`
	Password    string            `json:"password,omitempty"`
	Proxy       bool              `json:"proxy,omitempty"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Image       string            `json:"image,omitempty"`
	Rewrite     bool              `json:"rewrite,omitempty"`
	IOS         string            `json:"ios,omitempty"`
	Android     string            `json:"android,omitempty"`
	Geo         map[string]string `json:"geo,omitempty"`
}
type responseOptions struct {
	ID          string            `json:"id"`
	Domain      string            `json:"domain"`
	Key         string            `json:"key"`
	ExternalID  string            `json:"externalId"`
	URL         string            `json:"url"`
	Archived    bool              `json:"archived"`
	ExpiresAt   string            `json:"expiresAt"`
	ExpiredURL  string            `json:"expiredUrl"`
	Password    string            `json:"password"`
	Proxy       bool              `json:"proxy"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	Rewrite     bool              `json:"rewrite"`
	IOS         string            `json:"ios"`
	Android     string            `json:"android"`
	Geo         map[string]string `json:"geo"`
	PublicStats bool              `json:"publicStats"`
	TagID       string            `json:"tagId"`
	Tags        []Tag             `json:"tags"`
	Comments    string            `json:"comments"`
	ShortLink   string            `json:"shortLink"`
	QrCode      string            `json:"qrCode"`
	UtmSource   string            `json:"utm_source"`
	UtmMedium   string            `json:"utm_medium"`
	UtmCampaign string            `json:"utm_campaign"`
	UtmTerm     string            `json:"utm_term"`
	UtmContent  string            `json:"utm_content"`
	UserID      string            `json:"userId"`
	WorkspaceID string            `json:"workspaceId"`
	Clicks      int               `json:"clicks"`
	LastClicked string            `json:"lastClicked"`
	CreatedAt   string            `json:"createdAt"`
	UpdatedAt   string            `json:"updatedAt"`
	ProjectID   string            `json:"projectId"`
}
type Tag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type GetOptions struct {
	Domain     string `json:"domain"`
	Key        string `json:"key"`
	LinkId     string `json:"linkId"`
	ExternalId string `json:"externalId"`
}

type deleteResponse struct {
	Id string `json:"id"`
}
type sortBy string

const (
	CreatedAt   sortBy = "createdAt"
	Clicks      sortBy = "clicks"
	LastClicked sortBy = "lastClicked"
)

type GetListOptions struct {
	Domain       string   `json:"domain,omitempty"`
	TagIds       []string `json:"tagIds,omitempty"`
	TagNames     []string `json:"tagNames,omitempty"`
	Search       string   `json:"search,omitempty"`
	UserID       string   `json:"userId,omitempty"`
	ShowArchived bool     `json:"showArchived,omitempty"`
	WithTags     bool     `json:"withTags,omitempty"`
	Sort         sortBy   `json:"sort,omitempty"`
	Page         int      `json:"page,omitempty"`
}

type groupBy string

const (
	Domain groupBy = "domain"
	TagId  groupBy = "tagId"
)

type GetCountOptions struct {
	Domain       string   `json:"domain,omitempty"`
	TagIds       []string `json:"tagIds,omitempty"`
	TagNames     []string `json:"tagNames,omitempty"`
	Search       string   `json:"search,omitempty"`
	UserID       string   `json:"userId,omitempty"`
	ShowArchived bool     `json:"showArchived,omitempty"`
	WithTags     bool     `json:"withTags,omitempty"`
	GroupBy      groupBy  `json:"groupBy,omitempty"`
}
