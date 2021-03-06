package contentful

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"

	"evanjon.es/data"
)

type markdownProvider interface {
	Format(string) string
}

type contentful struct {
	md           markdownProvider
	space, token string
}

// contentful data type
type theirItem struct {
	Fields struct {
		Title     string
		Slug      string
		Img       string
		ImgAlt    string
		Alt       string
		Desc      string
		ShortDesc string
		Raw       template.HTML
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

func Default(md markdownProvider, space, token string) contentful {
	return contentful{md, space, token}
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
		return ret, errors.New("Hm... Not sure I can find what you're looking for. What was it called again?")
	}

	ret = item{
		entry.Fields.Title,
		c.md.Format(entry.Fields.Desc),
		entry.Fields.ShortDesc,
		entry.Fields.Img,
		entry.Fields.ImgAlt,
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
		return ret, errors.New("Nothing here. Darn, I know we have posts. Where could they be?! Hold on. I'll go look in the back.")
	}

	for _, single := range entry.Items {
		ret = append(ret, item{
			single.Fields.Title,
			c.md.Format(single.Fields.Desc),
			single.Fields.ShortDesc,
			single.Fields.Img,
			single.Fields.ImgAlt,
			single.Fields.Slug,
			single.Sys.ID,
		})
	}

	return ret, nil
}
