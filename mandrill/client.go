package mandrill

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ENDPOINT string = "https://mandrillapp.com/api/1.0/%s.json"

type Client struct {
	key string
}

type ErrorResult struct {
	Status  SendStatus `json:"status"`
	Code    int        `json:"code"`
	Name    string     `json:"name"`
	Message string     `json:"message"`
}

func (e *ErrorResult) Error() string {
	return e.Message
}

func (c *Client) Call(endpoint string, data map[string]interface{}, results interface{}) error {
	url := fmt.Sprintf(ENDPOINT, endpoint)

	if data == nil {
		data = make(map[string]interface{})
	}
	data["key"] = c.key

	// wouldn't it be more efficient to json.Encode directly into a buffer?
	j, _ := json.Marshal(data)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

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

	if response.StatusCode >= 400 {
		e := &ErrorResult{}
		err := json.Unmarshal(body, e)
		if err != nil {
			return fmt.Errorf("Unknown client error: %s", response.Status)
		}
		return e
	}

	return json.Unmarshal(body, results)
}

func NewClient(key string) *Client {
	return &Client{key}
}
