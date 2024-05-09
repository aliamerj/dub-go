package workspace

type Options struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Domain string `json:"domain,omitempty"`
}

type responseWorkspace struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Slug              string   `json:"slug"`
	Logo              *string  `json:"logo"`
	Usage             int      `json:"usage"`
	UsageLimit        int      `json:"usageLimit"`
	LinksUsage        int      `json:"linksUsage"`
	LinksLimit        int      `json:"linksLimit"`
	DomainsLimit      int      `json:"domainsLimit"`
	TagsLimit         int      `json:"tagsLimit"`
	UsersLimit        int      `json:"usersLimit"`
	Plan              string   `json:"plan"`
	StripeId          *string  `json:"stripeId"`
	BillingCycleStart int      `json:"billingCycleStart"`
	CreatedAt         string   `json:"createdAt"`
	Users             []user   `json:"users"`
	Domains           []domain `json:"domains"`
	InviteCode        *string  `json:"inviteCode"`
}

type user struct {
	Role string `json:"role"`
}
type domain struct {
	Slug    string `json:"slug"`
	Primary bool   `json:"primary"`
}
