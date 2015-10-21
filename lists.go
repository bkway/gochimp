package gochimp

import(
//	"time"
//	"fmt"
//	"log"
)

// FIXME merge with Mandrill Message
type Message struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject   string `json:"subject"`
	Language  string `json:"language"`
}

type VisibilityStatus string
const(
	PublicVisibility    VisibilityStatus   = "pub"
	PrivateVisibility                      = "prv"
)

type SubscriptionStatus string
const(
	Subscribed          SubscriptionStatus = "subscribed"
	Unsubscribed                           = "unsubscribed"
	SubscriptionCleaned                    = "cleaned"
	SubscriptionPending                    = "pending"
)

type List struct {
	Id                            string           `json:"id,omitempty"`
	Name                          string           `json:"name"`
	Contact                       Contact          `json:"contact"`
	PermissionReminder            string           `json:"permission_reminder"`
	UseArchiveBar                 bool             `json:"use_archive_bar,omitempty"`
	CampaignDefaults struct {
		FromName                  string           `json:"from_name"`
		FromEmail                 string           `json:"from_email"`
		Subject                   string           `json:"subject"`
		Language                  string           `json:"language"`
	}                                              `json:"campaign_defaults"`
	NotifyOnSubscribe             string           `json:"notify_on_subscribe,omitempty"`
	NotifyOnUnsubscribe           string           `json:"notify_on_unsubscribe,omitempty"`
	DateCreated                   string           `json:"date_created,omitempty"`
	ListRating                    uint             `json:"list_rating,omitempty"`
	EmailTypeOption               bool             `json:"email_type_option"`
	SubscribeUrlShort             string           `json:"subscribe_url_short,omitempty"`
	SubscribeUrlLong              string           `json:"subscribe_url_long,omitempty"`
	BeamerAddress                 string           `json:"beamer_address,omitempty"`
	Visibility                    VisibilityStatus `json:"visibility,omitempty"`
	Modules                       []string         `json:"modules,omitempty"`
	Stats struct {
		MemberCount               uint64           `json:"member_count"`
		UnsubscribeCount          uint64           `json:"unsubscribe_count"`
		CleanedCount              uint64           `json:"cleaned_count"`
		MemberCountSinceSend      uint64           `json:"member_count_since_send"`
		UnsubscribeCountSinceSend uint64           `json:"unsubscribe_count_since_send"`
		CleanedCountSinceSend     uint64           `json:"cleaned_count_since_send"`
		CampaignCount             uint64           `json:"campaign_count"`

		CampaignLastSent          string           `json:"campaign_last_sent,omitempty"`

		MergeFieldCount           int              `json:"merge_field_count"`

		AvgSubRate                float32          `json:"avg_sub_rate"`
		AvgUnsubRate              float32          `json:"avg_unsub_rate"`
		TargetSubRate             float32          `json:"target_sub_rate"`
		OpenRate                  float32          `json:"open_rate"`
		ClickRate                 float32          `json:"click_rate"`

		LastSubDate               string           `json:"last_sub_date,omitempty"`
		LastUnsubDate             string           `json:"last_unsub_date,omitempty"`
	}                                              `json:"stats,omitempty"`
}

func (c *Client) Lists() ([]List, error) {
	// GET /lists/
	var results struct {
		Lists []List `json:"lists"`
	}
	err := c.Call("GET", "lists", nil, &results)
	if err == nil {
		return results.Lists, nil
	}
	return nil, err
}

func (c *Client) CreateList(list *List) (List, error) {
	// POST /lists/
	var results List
	err := c.Call("POST", "lists", list, &results)
	return results, err
}

func (c *Client) List(listId string) (List, error) {
	// GET /lists/{list_id}
	var results List
	err := c.Call("GET", "lists/" + listId, nil, &results)
	return results, err
}

func (c *Client) UpdateList(list *List) (List, error) {
	// PATCH /lists/{list_id}
	var results List
	err := c.Call("PATCH", "lists/" + list.Id, list, &results)
	return results, err
}

func (c *Client) DeleteList(listId string) error {
	// DELETE /lists/{list_id}
	var results interface{}
	return c.Call("DELETE", "lists/" + listId, nil, &results)
}
