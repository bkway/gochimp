package go_mailchimp

import(
	"github.com/jordan-wright/email"
//	"net/smtp"
	"encoding/json"
	"strings"
	"time"
	"strconv"
)

type Flag string
const(
	Autotext Flag      = "Autotext"
	AutoHtml           = "AutoHtml"
	URLStripQS         = "URLStripQS"
	PreserveRecipients = "PreserveRecipients"
	InlineCSS          = "InlineCSS"
	ViewContentLink    = "ViewContentLink"
	Important          = "Important"
)

const(
	Host string = "smtp.mandrillapp.com"
	Port        = "587"
)

type MergeLang string
const(
	Mailchimp MergeLang    = "mailchimp"
	Handlebars             = "handlebars"
)

type Google struct {
	Domains  []string
	Campaign string
}

type Template struct {
	Name               string
	MergeLanguage      MergeLang
	Tracks             string
	TrackingDomain     string
	GoogleAnalytics    *Google
	SigningDomain      string
	ReturnPathDomain   string
	Subaccount         string
	BccAddress         string
	IpPool             string
	Flags              map[Flag]bool
	//	MergeVars          interface{}
	//	Metadata           *map[string]string
	//	Tags               *[]string
	//	SendAt             *time.Time
}

func NewTemplate(name string, lang MergeLang) *Template {
	return &Template{
		Name: name,
		MergeLanguage: lang,
		Flags: make(map[Flag]bool),
	}
}

func (p *Template) AddGoogleAnalytics(campaign string, domains []string) *Template {
	p.GoogleAnalytics = &Google{
		Campaign: campaign,
		Domains: domains,
	}
	return p
}

func (p *Template) Track(opens, htmlClicks, textClicks bool) *Template {
	tracks := make([]string, 0, 0)
	if opens {
		tracks = append(tracks, "opens")
	}
	if htmlClicks && textClicks {
		tracks = append(tracks, "clicks_all")
	} else if htmlClicks {
		tracks = append(tracks, "clicks_htmlonly")
	} else if textClicks {
		tracks = append(tracks, "clicks_textonly")
	}
	p.Tracks = strings.Join(tracks, ",")
	return p
}

func (p *Template) Render(data interface{}, tags []string, meta map[string]string, at *time.Time) (*email.Email, error) {
	out := email.NewEmail()
	out.Headers.Set("X-MC-Template", p.Name)

	j, err := json.Marshal(data)
	if err != nil {
		return out, err
	}
	out.Headers.Set("X-MC-MergeVars", string(j))

	if tags != nil && len(tags) > 0 {
		out.Headers.Set("X-MC-Tags", strings.Join(tags, ","))
	}

	if meta != nil {
		j, err := json.Marshal(meta)
		if err == nil {
			out.Headers.Set("X-MC-Meta", string(j))
		}
	}

	if at != nil {
		// handle time
	}

	// common

	if p.Tracks != "" {
		out.Headers.Set("X-MC-Tracks", p.Tracks)
	}

	if p.TrackingDomain != "" {
		out.Headers.Set("X-MC-TrackingDomain", p.TrackingDomain)
	}

	if p.GoogleAnalytics != nil {
		if p.GoogleAnalytics.Campaign != "" {
			out.Headers.Set("X-MC-GoogleAnalyticsCampaign", p.GoogleAnalytics.Campaign)
		}
		if len(p.GoogleAnalytics.Domains) > 0 {
			out.Headers.Set("X-MC-GoogleAnalytics", strings.Join(p.GoogleAnalytics.Domains, ","))
		}
	}

	if p.SigningDomain != "" {
		out.Headers.Set("X-MC-SigningDomain", p.SigningDomain)
	}

	if p.ReturnPathDomain != "" {
		out.Headers.Set("X-MC-ReturnPathDomain", p.ReturnPathDomain)
	}

	if p.Subaccount != "" {
		out.Headers.Set("X-MC-Subaccount", p.Subaccount)
	}

	if p.BccAddress != "" {
		out.Headers.Set("X-MC-BccAddress", p.BccAddress)
	}

	if p.IpPool != "" {
		out.Headers.Set("X-MC-IpPool", p.IpPool)
	}

	for flag, val := range p.Flags {
		out.Headers.Set("X-MC-" + string(flag), strconv.FormatBool(val))
	}

	return out, nil
}
