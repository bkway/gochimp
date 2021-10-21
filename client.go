package gochimp

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const ENDPOINT string = "https://%s.api.mailchimp.com/3.0/%s"

type Client struct {
	dc  string
	key string
}

type Contact struct {
	Company  string `json:"company"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitempty"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
	Phone    string `json:"phone,omitempty"`
}

type Detail struct {
	AccountId        string  `json:"account_id"`
	AccountName      string  `json:"account_name"`
	Contact          Contact `json:"contact"`
	LastLogin        string  `json:"last_login"`
	TotalSubscribers int64   `json:"total_subscribers"`
}

func EmailToHash(email string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(email)))
}

func NewClient(key string) *Client {
	dc := strings.SplitN(key, "-", 2)[1]
	return &Client{dc, key}
}

func (c *Client) Call(method string, endpoint string, data interface{}, results interface{}) error {
	url := fmt.Sprintf(ENDPOINT, c.dc, endpoint)

	buf := new(bytes.Buffer)

	if data != nil {
		// wouldn't it be more efficient to json.Encode directly into a buffer?
		j, err := json.Marshal(data)
		if err != nil {
			return err
		}

		pretty, _ := json.MarshalIndent(data, "", "\t")
		log.Println(string(pretty))

		buf.Write(j)
	}

	request, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("GoLangAPIClient", c.key)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, results)
}

func (c *Client) Details() (Detail, error) {
	var result Detail
	err := c.Call("GET", "", nil, &result)
	return result, err
}
