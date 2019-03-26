package contentful

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"html/template"
	"errors"
)

// Item - todo
type Item struct {
	Fields struct {
		Title string
		Slug  string
		Image string
		Alt string
		Desc string
		Raw	template.HTML
	}
	Sys struct {
		ID string
	}
}

// ItemList - todo
type ItemList struct {
	Items []Item
}

// Service - todo
type Service struct {
	BlogAPI func() string
	PostAPI func(post string) string
}

// Default - todo
func Default(space string, token string) *Service {
	return &Service{
		BlogAPI: func() string {
			return "https://cdn.contentful.com/spaces/" + space + "/entries?access_token=" + token + "&content_type=post"
		},
		PostAPI: func(post string) string {
			return "https://cdn.contentful.com/spaces/" + space + "/entries/" + post + "?access_token=" + token
		},
	}
}

// GetPosts - todo
func (c *Service) GetPosts() ([]Item, error) {
	req, err := http.Get(c.BlogAPI())

	defer req.Body.Close()

	if err != nil {
		return ItemList{}.Items, err
	}

	res, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return ItemList{}.Items, err
	}

	var list ItemList
	err = json.Unmarshal(res, &list)

	if err != nil {
		return ItemList{}.Items, err
	}

	if len(list.Items) < 1 {
		return ItemList{}.Items, errors.New("NotFound")
	}

	return list.Items, nil
}

// GetPost - todo
func (c *Service) GetPost(post string) (Item, error) {
	req, err := http.Get(c.PostAPI(post))

	defer req.Body.Close()
	
	if err != nil {
		return Item{}, err
	}

	res, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return Item{}, err
	}

	var item Item
	err = json.Unmarshal(res, &item)

	if err != nil {
		return Item{}, err
	}

	item.Fields.Raw = template.HTML(item.Fields.Desc)

	if item.Sys.ID == "NotFound" {
		return item, errors.New(item.Sys.ID)
	}

	return item, nil
}
