package mandrill

/*
import(
	"time"
)

type StatWindow string
const(
	Today      StatWindow = "today"
	Last7Days             = "last_7_days"
	Last30Days            = "last_30_days"
	Last60Days            = "last_60_days"
	Last90Days            = "last_90_days"
	AllTime               = "all_time"
)

type Stat struct {
	Address      email.Email          `json:"address,omitempty"`
	CreatedAt    *time.Time           `json:"created_at,omitempty"`
	Sent         uint                 `json:"sent"`
	HardBounces  uint                 `json:"hard_bounces"`
	SoftBounces  uint                 `json:"soft_bounces"`
	Rejects      uint                 `json:"rejects"`
	Complaints   uint                 `json:"complaints"`
	Unsubs       uint                 `json:"unsubs"`
	Opens        uint                 `json:"opens"`
	UniqueOpens  uint                 `json:"unique_opens"`
	Clicks       uint                 `json:"clicks"`
	UniqueClicks uint                 `json:"unique_clicks"`
}

type Info struct {
	Username     string               `json:"username"`
	CreatedAt    *time.Time           `json:"created_at"`
	PublicId     string               `json:"public_id"`
	Reputation   uint8                `json:"reputation"`
	HourlyQuota  uint                 `json:"hourly_quota"`
	Backlog      uint                 `json:"backlog"`
	Stats        map[StatWindow]Stat  `json:"stats"`
}

func (p *Client) Info() (Info, error) {
	var info Info
	err := p.call("users/info", nil, &info)
	return info, err
}
*/

func (p *Client) Ping() (bool, error) {
	var pong string
	err := p.Call("users/ping", nil, &pong)
	return pong == "PONG!", err
}

/*
func (p *Client) Ping2() (bool, error) {
	var pong struct{ Ping string `json:"PING"` }
	err :=  p.call("users/ping2", nil, &pong)
	if err == nil {
		return pong.Ping == "PONG!", nil
	}
	return false, err
}

func (p *Client) Senders() ([]Stat, error) {
	var senders []Stat
	err := p.call("users/senders", nil, &senders)
	return senders, err
}
*/
