package gochimp

import(
	"fmt"
)

type EmailType string
const(
	HtmlEmail EmailType = "html"
	TextEmail           = "text"
)

type Member struct {
	Id               string                 `json:"id,omitempty"`
	EmailAddress     string                 `json:"email_address"`
	UniqueEmailId    string                 `json:"unique_email_id,omitempty"`
	EmailType        EmailType              `json:"email_type,omitempty"`
	Status           SubscriptionStatus     `json:"status,omitempty"`
	StatusIfNew      SubscriptionStatus     `json:"status_if_new,omitempty"`
	MergeFields      map[string]interface{} `json:"merge_fields,omitempty"`
	Interests        map[string]bool        `json:"interests,omitempty"`
	Stats struct {
		AvgOpenRate  float32                `json:"avg_open_rate,omitempty"`
		AvgClickRate float32                `json:"avg_click_rate,omitempty"`
	}                                       `json:"stats,omitempty"`
	IpSignup         string                 `json:"ip_signup,omitempty"`
	TimestampSignup  string                 `json:"timestamp_signup,omitempty"`
	IpOpt            string                 `json:"ip_opt,omitempty"`
	TimestampOpt     string                 `json:"timestamp_opt,omitempty"`
	MemberRating     uint                   `json:"member_rating,omitempty"`
	LastChanged      string                 `json:"last_changed,omitempty"`
	Language         string                 `json:"language,omitempty"`
	Vip              bool                   `json:"vip,omitempty"`
	EmailClient      string                 `json:"email_client,omitempty"`
	Location struct {
		Lat          float64                `json:"latitude,omitempty"`
		Lon          float64                `json:"longitude,omitempty"`
		Gmtoff       int                    `json:"gmtoff,omitempty"`
		Dstoff       int                    `json:"dstoff,omitempty"`
		CountryCode  string                 `json:"country_code,omitempty"`
		Timezone     string                 `json:"timezone,omitempty"`
	}                                       `json:"location,omitempty"`
	LastNote struct {
		Id           string                 `json:"note_id,omitempty"`
		CreatedAt    string                 `json:"created_at,omitempty"`
		CreatedBy    string                 `json:"created_by,omitempty"`
		Note         string                 `json:"note,omitempty"`
	}                                       `json:"last_note,omitempty"`
	ListId           string                 `json:"list_id,omitempty"`
}

func (m *Member) SetInterest(key string, val bool) *Member {
	if m.Interests == nil {
		m.Interests = make(map[string]bool)
	}
	m.Interests[key] = val
	return m
}

func (m *Member) SetMergeField(key string, val interface{}) *Member {
	if m.MergeFields == nil {
		m.MergeFields = make(map[string]interface{})
	}
	m.MergeFields[key] = val
	return m
}

func NewMember(email string) *Member {
	return &Member{ Id: EmailToHash(email), EmailAddress: email }
}

func (c *Client) Members(listId string) ([]*Member, error) {
	// GET /lists/{list_id}/members
	var results struct {
			Members []*Member `json:"members"`
		}
	err := c.Call("GET", fmt.Sprintf("lists/%s/members", listId), nil, &results)
	if err == nil {
		return results.Members, nil
	}
	return nil, err
}

func (c *Client) AddMember(listId string, member *Member) (*Member, error) {
	// POST /lists/{list_id}/members
	err := c.Call("POST", fmt.Sprintf("lists/%s/members", listId), member, member)
	return member, err
}

func (c *Client) Member(listId, emailHash string) (*Member, error) {
	// GET /lists/{list_id}/members/{id}
	results := &Member{}
	err := c.Call("GET", fmt.Sprintf("lists/%s/members/%s", listId, emailHash), nil, results)
	return results, err
}

func (c *Client) UpsertMember(listId string, member *Member) (*Member, error) {
	// PUT /lists/{list_id}/members/{id}
	err := c.Call("PUT", fmt.Sprintf("lists/%s/members/%s", listId, member.Id), member, member)
	return member, err
}

func (c *Client) UpdateMember(listId string, member *Member) (*Member, error) {
	// PATCH /lists/{list_id}/members/{id}
	err := c.Call("PATCH", fmt.Sprintf("lists/%s/members/%s", listId, member.Id), member, member)
	return member, err
}

func (c *Client) DeleteMember(listId, emailHash string) error {
	// DELETE /lists/{list_id}/members/{id}
	var results interface{}
	return c.Call("DELETE", fmt.Sprintf("lists/%s/members/%s", listId, emailHash), nil, &results)
}
