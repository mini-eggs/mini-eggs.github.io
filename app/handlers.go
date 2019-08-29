package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app app) ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (app app) page(c *gin.Context, id string) {
	item, err := app.c.Get(id)
	if err != nil {
		app.error(c, err)
		return
	}
	c.HTML(http.StatusOK, "item.html", item)
}

func (app app) post(c *gin.Context) {
	item, err := app.c.Get(c.Param("id"))
	if err != nil {
		app.error(c, err)
		return
	}
	c.HTML(http.StatusOK, "item.html", item)
}

func (app app) posts(c *gin.Context) {
	items, err := app.c.List("post")
	if err != nil {
		app.error(c, err)
		return
	}
	c.HTML(http.StatusOK, "list.html", items)
}

func (app app) rss(c *gin.Context) {
	list, err := app.c.List("post")
	if err != nil {
		app.error(c, err)
		return
	}

	rss, err := app.r.List(list)
	if err != nil {
		app.error(c, err)
		return
	}

	c.String(http.StatusOK, rss)
}

func (app app) error(c *gin.Context, err error) {
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{"MSG": err.Error()})
}

func (app app) notFound(c *gin.Context) {
	err := errors.New("404, not found. :/")
	c.HTML(http.StatusNotFound, "error.html", gin.H{"MSG": err.Error()})
}
