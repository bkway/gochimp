package mandrill

type SendType string

const (
	SendTo  SendType = "to"
	SendCC           = "cc"
	SendBCC          = "bcc"
)

type Recipient struct {
	Email     string            `json:"email"`
	Name      string            `json:"name,omitempty"`
	Type      SendType          `json:"type,omitempty"`
	MergeVars []Variable        `json:"-"`
	Metadata  map[string]string `json:"-"`
}

type MergeLanguage string

const (
	Mailchimp  MergeLanguage = "mailchimp"
	Handlebars               = "handlebars"
)

type RcptMergeVars struct {
	Rcpt string     `json:"rcpt"`
	Vars []Variable `json:"vars"`
}

type RcptMetadata struct {
	Rcpt   string            `json:"rcpt"`
	Values map[string]string `json:"values"`
}

type Attachment struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Message struct {
	Html string `json:"html,omitempty"`
	Text string `json:"text,omitempty"`

	Subject   string `json:"subject,omitempty"`
	FromEmail string `json:"from_email,omitempty"`
	FromName  string `json:"from_name,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`

	Important          bool `json:"important,omitempty"`
	TrackOpens         bool `json:"track_opens,omitempty"`
	TrackClicks        bool `json:"track_clicks,omitempty"`
	AutoText           bool `json:"auto_text,omitempty"`
	AutoHtml           bool `json:"auto_html,omitempty"`
	InlineCSS          bool `json:"inline_css,omitempty"`
	UrlStripQs         bool `json:"url_strip_qs,omitempty"`
	PreserveRecipients bool `json:"preserve_recipients,omitempty"`
	ViewContentLink    bool `json:"view_content_link,omitempty"`

	BCCAddress       string        `json:"bcc_address,omitempty"`
	TrackingDomain   string        `json:"tracking_domain,omitempty"`
	SigningDomain    string        `json:"signing_domain,omitempty"`
	ReturnPathDomain string        `json:"return_path_domain,omitempty"`
	Merge            bool          `json:"merge,omitempty"`
	MergeLanguage    MergeLanguage `json:"merge_language,omitempty"`

	GlobalMergeVars []Variable `json:"global_merge_vars,omitempty"`

	Tags                    []string       `json:"tags,omitempty"`
	SubAccount              string         `json:"sub_account,omitempty"`
	GoogleAnalyticsDomains  []string       `json:"google_analytics_domains,omitempty"`
	GoogleAnalyticsCampaign []string       `json:"google_analytics_campaign,omitempty"`
	Metadata                []RcptMetadata `json:"metadata,omitempty"`

	Recipients []Recipient `json:"to,omitempty"`
	// built from Recipients
	MergeVars         []RcptMergeVars `json:"merge_vars,omitempty"`
	RecipientMetadata []RcptMetadata  `json:"recipient_metadata,omitempty"`

	Attachments []Attachment `json:"attachments,omitempty"`
	Images      []Attachment `json:"images,omitempty"`
}

func NewMessage() *Message {
	msg := Message{}
	return &msg
}

func (m *Message) AddRecipient(t SendType, r Recipient) *Message {
	r.Type = t
	m.Recipients = append(m.Recipients, r)

	if r.MergeVars != nil {
		m.MergeVars = append(m.MergeVars, RcptMergeVars{Rcpt: r.Email, Vars: r.MergeVars})
	}
	if r.Metadata != nil {
		m.Metadata = append(m.Metadata, RcptMetadata{Rcpt: r.Email, Values: r.Metadata})
	}
	return m
}

func (m *Message) To(email string, name string) *Message {
	return m.AddRecipient(SendTo, Recipient{Email: email, Name: name})
}

func (m *Message) CC(email string, name string) *Message {
	return m.AddRecipient(SendCC, Recipient{Email: email, Name: name})
}

func (m *Message) BCC(email string, name string) *Message {
	return m.AddRecipient(SendBCC, Recipient{Email: email, Name: name})
}

func (m *Message) WithValue(key string, content interface{}) *Message {
	m.GlobalMergeVars = append(m.GlobalMergeVars, Variable{key, content})
	return m
}

func (m *Message) WithValues(vals map[string]interface{}) *Message {
	for k, v := range vals {
		m.WithValue(k, v)
	}
	return m
}

type SendStatus string

const (
	MessageSent      SendStatus = "sent"
	MessageQueued               = "queued"
	MessageScheduled            = "scheduled"
	MessageRejected             = "rejected"
	MessageInvalid              = "invalid"
	ClientError                 = "error"
)

type RejectType string

const (
	RejectedHardBounce    RejectType = "hard-bounce"
	RejectedSoftBounce               = "soft-bounce"
	RejectedSpam                     = "spam"
	RejectedUnsub                    = "unsub"
	RejectedCustom                   = "custom"
	RejectedInvalidSender            = "invalid-sender"
	RejectedInvalid                  = "invalid"
	RejectedModeLimit                = "test-mode-limit"
	RejectedRule                     = "rule"
)

type SendResult struct {
	Email        string     `json:"email"`
	Status       SendStatus `json:"status"`
	RejectReason RejectType `json:"reject_reason"`
	Id           string     `json:"id"`
}

type SendOpt func(map[string]interface{})

func SendAt(val string) SendOpt {
	return func(vals map[string]interface{}) {
		vals["send_at"] = val
	}
}

func IPPool(val string) SendOpt {
	return func(vals map[string]interface{}) {
		vals["ip_pool"] = val
	}
}

func Async(val bool) SendOpt {
	return func(vals map[string]interface{}) {
		vals["async"] = val
	}
}

func prepMessageSend(msg *Message, opts ...SendOpt) map[string]interface{} {
	request := map[string]interface{}{"message": msg}

	if opts != nil {
		for _, opt := range opts {
			opt(request)
		}
	}
	return request
}

func (p *Client) MessagesSend(msg *Message, opts ...SendOpt) ([]SendResult, error) {
	request := prepMessageSend(msg, opts...)

	var result []SendResult
	err := p.Call("messages/send", request, &result)
	return result, err
}

func (p *Client) MessagesSendTemplate(name string, content []Variable, msg *Message, opts ...SendOpt) ([]SendResult, error) {
	request := prepMessageSend(msg, opts...)

	request["template_name"] = name
	request["template_content"] = content

	var result []SendResult
	err := p.Call("messages/send-template", request, &result)
	return result, err
}

/*
func (p *Client) MessagesSearch() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesSearchTimeSeries() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesInfo() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesContent() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesParse() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesSendRaw() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesListScheduled() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesCancelScheduled() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

func (p *Client) MessagesReschedule() (interface{}, error) {
	return p.Call("messages", "info", nil)
}

*/
