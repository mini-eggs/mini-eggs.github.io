package contentful

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"

	"evanjon.es/data"
)

type contentful struct {
	space, token string
}

// contentful data type
type theirItem struct {
	Fields struct {
		Title string
		Slug  string
		Img   string
		Alt   string
		Desc  string
		Raw   template.HTML
	}
	Sys struct {
		ID string
	}
}

// contentful data type
type theirList struct {
	Items []theirItem
}

// contentful service

func Default(space, token string) contentful {
	return contentful{space, token}
}

// url construction

func (c contentful) postURL(id string) string {
	return "https://cdn.contentful.com/spaces/" + c.space + "/entries/" + id + "?access_token=" + c.token
}

func (c contentful) listURL(id string) string {
	return "https://cdn.contentful.com/spaces/" + c.space + "/entries?access_token=" + c.token + "&content_type=" + id + "&order=-sys.createdAt"
}

// public request/get methods

func (c contentful) Get(id string) (ret data.Item, err error) {
	req, err := http.Get(c.postURL(id))
	if err != nil {
		return
	}
	defer req.Body.Close()

	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	var entry theirItem
	err = json.Unmarshal(res, &entry)
	if err != nil {
		return
	}

	entry.Fields.Raw = template.HTML(entry.Fields.Desc)

	if entry.Sys.ID == "NotFound" {
		err = errors.New(entry.Sys.ID)
		return
	}

	ret = item{
		entry.Fields.Title,
		entry.Fields.Desc,
		entry.Fields.Img,
		entry.Fields.Slug,
		entry.Sys.ID,
	}

	return ret, nil
}

func (c contentful) List(id string) (ret []data.Item, err error) {
	req, err := http.Get(c.listURL(id))
	if err != nil {
		return
	}
	defer req.Body.Close()

	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	var entry theirList
	err = json.Unmarshal(res, &entry)
	if err != nil {
		return
	}

	if len(entry.Items) < 1 {
		err = errors.New("failed to find any posts")
		return
	}

	for _, single := range entry.Items {
		ret = append(ret, item{
			single.Fields.Title,
			single.Fields.Desc,
			single.Fields.Img,
			single.Fields.Slug,
			single.Sys.ID,
		})
	}

	return ret, nil
}
