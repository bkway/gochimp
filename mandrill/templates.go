package mandrill

type Template struct {
	Slug   string   `json:"slug,omitempty"`
	Name   string   `json:"name"`
	Labels []string `json:"labels,omitempty"`

	Code string `json:"code,omitempty"`

	Subject   string `json:"subject,omitempty"`
	FromEmail string `json:"from_email,omitempty"`
	FromName  string `json:"from_name,omitempty"`

	Text string `json:"text,omitempty"`

	PublishCode      string `json:"publish_code,omitempty"`
	PublishSubject   string `json:"publish_subject,omitempty"`
	PublishFromEmail string `json:"publish_from_email,omitempty"`
	PublishFromName  string `json:"publish_from_name,omitempty"`
	PublishText      string `json:"publish_text,omitempty"`

	PublishedAt string `json:"published_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	Publish     bool   `json:"publish,omitempty"`
}

/*
func (p *Client) TemplatesAdd() (interface{}, error) {
	return p.call("user", "info", nil)
}
*/

func (p *Client) TemplatesInfo(name string) (Template, error) {
	request := map[string]interface{}{"name": name}
	var result Template
	err := p.Call("templates/info", request, &result)
	return result, err
}

/*
func (p *Client) TemplatesUpdate() (interface{}, error) {
	return p.call("user", "info", nil)
}

func (p *Client) TemplatesPublish() (interface{}, error) {
	return p.call("user", "info", nil)
}

func (p *Client) TemplatesDelete() (interface{}, error) {
	return p.call("user", "info", nil)
}
*/

func (p *Client) TemplatesList(label string) ([]Template, error) {
	request := make(map[string]interface{})
	if label != "" {
		request["label"] = label
	}
	var result []Template
	err := p.Call("templates/list", request, &result)
	return result, err
}

/*
func (p *Client) TemplatesTimeSeries() (interface{}, error) {
	return p.call("user", "info", nil)
}
*/
func (p *Client) TemplatesRender(name string, content []Variable, vars []Variable) (interface{}, error) {
	request := map[string]interface{}{
		"template_name":    name,
		"template_content": content,
	}
	if vars != nil {
		request["merge_vars"] = vars
	}
	var result map[string]string
	err := p.Call("templates/render", request, &result)
	return result["html"], err
}
