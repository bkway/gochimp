package gochimp

import (
	"fmt"
)

type CategoryType string

const (
	Checkboxes CategoryType = "checkboxes"
	Dropdown                = "dropdown"
	Radio                   = "radio"
	Hidden                  = "hidden"
)

type Category struct {
	Id           string       `json:"id"`
	ListId       string       `json:"list_id"`
	Title        string       `json:"title"`
	DisplayOrder uint         `json:"display_order"`
	Type         CategoryType `json:"type"`
}

type Interest struct {
	Id           string `json:"id"`
	ListId       string `json:"list_id"`
	CategoryId   string `json:"category_id"`
	Name         string `json:"name"`
	DisplayOrder uint   `json:"display_order"`
}

func (c *Client) InterestCategories(listId string) ([]Category, error) {
	// GET /lists/{list_id}/interest-categories
	var results struct {
		Categories []Category `json:"categories"`
	}
	err := c.Call("GET", fmt.Sprintf("/lists/%s/interest-categories", listId), nil, &results)
	if err == nil {
		return results.Categories, nil
	}
	return nil, err
}

func (c *Client) CreateInterestCategory(listId string, category *Category) (*Category, error) {
	// POST /lists/{list_id}/interest-categories
	err := c.Call("POST", fmt.Sprintf("/lists/%s/interest-categories", listId), category, category)
	return category, err
}

func (c *Client) InterestCategory(listId, categoryId string) (Category, error) {
	// GET /lists/{list_id}/interest-categories/{id}
	var result Category
	err := c.Call("GET", fmt.Sprintf("/lists/%s/interest-categories/%s", listId, categoryId), nil, &result)
	return result, err
}

func (c *Client) UpdateInterestCategory(listId, categoryId string, category *Category) (*Category, error) {
	// PATCH /lists/{list_id}/interest-categories/{id}
	err := c.Call("PATCH", fmt.Sprintf("/lists/%s/interest-categories/%s", listId, categoryId), category, category)
	return category, err
}

func (c *Client) DeleteInterestCategory(listId, categoryId string) error {
	// DELETE /lists/{list_id}/interest-categories/{id}
	var result interface{}
	return c.Call("DELETE", fmt.Sprintf("/lists/%s/interest-categories/%s", listId, categoryId), nil, result)
}

func (c *Client) Interests(listId, categoryId string) ([]Interest, error) {
	// GET /lists/{list_id}/interest-categories/{category_id}/interests
	var result struct {
		Interests []Interest `json:"interests"`
	}
	err := c.Call("GET", fmt.Sprintf("/lists/%s/interest-categories/%s/interests", listId, categoryId), nil, &result)
	return result.Interests, err
}

func (c *Client) AddInterest(listId, categoryId, interest *Interest) (*Interest, error) {
	// POST /lists/{list_id}/interest-categories/{category_id}/interests
	err := c.Call("POST", fmt.Sprintf("/lists/%s/interest-categories/%s/interests", listId, categoryId), interest, interest)
	return interest, err
}

func (c *Client) Interest(listId, categoryId, id string) (Interest, error) {
	// GET /lists/{list_id}/interest-categories/{category_id}/interests/{id}
	var results Interest
	err := c.Call("GET", fmt.Sprintf("/lists/%s/interest-categories/%s/interests/%s", listId, categoryId, id), nil, &results)
	return results, err
}

func (c *Client) UpdateInterest(listId, categoryId, interest *Interest) (*Interest, error) {
	// PATCH /lists/{list_id}/interest-categories/{category_id}/interests/{id}
	err := c.Call("PATCH", fmt.Sprintf("/lists/%s/interest-categories/%s/interests/%s", listId, categoryId, interest.Id), interest, interest)
	return interest, err
}

func (c *Client) DeleteInterest(listId, categoryId, id string) error {
	// DELETE /lists/{list_id}/interest-categories/{category_id}/interests/{id}
	var results interface{}
	return c.Call("DELETE", fmt.Sprintf("/lists/%s/interest-categories/%s/interests/%s", listId, categoryId, id), nil, &results)
}
