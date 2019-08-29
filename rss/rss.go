package rss

import (
	"evanjon.es/data"
	"github.com/gorilla/feeds"
)

// Basically a wrapper over "github.com/gorilla/feeds" to fit our
// data.Item data type.

type rss struct {
	title, desc, link, author, email string
}

func Default(title, desc, link, author, email string) rss {
	return rss{title, desc, link, author, email}
}

func (r rss) rss(item data.Item) *feeds.Item {
	return &feeds.Item{
		Title:       item.Title(),
		Description: item.Desc(),
		Author:      &feeds.Author{Name: r.author, Email: r.email},
		Link:        &feeds.Link{Href: r.link + "post/" + item.Slug() + "/" + item.ID()},
	}
}

func (r rss) List(items []data.Item) (string, error) {
	feed := &feeds.Feed{
		Title:       r.title,
		Link:        &feeds.Link{Href: r.link},
		Description: r.desc,
		Author:      &feeds.Author{Name: r.author, Email: r.email},
	}

	for _, item := range items {
		feed.Items = append(feed.Items, r.rss(item))
	}

	return feed.ToRss()
}
