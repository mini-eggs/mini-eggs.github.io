package app

import (
	"errors"
	"net/http"

	"evanjon.es/data"
	"github.com/gin-gonic/gin"
)

type contentProvider interface {
	Get(string) (data.Item, error)
	List(string) ([]data.Item, error)
}

type rssProvider interface {
	List([]data.Item) (string, error)
}

type app struct {
	s *gin.Engine
	c contentProvider
	r rssProvider
}

func Default(cp contentProvider, rss rssProvider) app {
	a := app{gin.Default(), cp, rss}

	a.s.Static("/static", "./static")
	a.s.LoadHTMLGlob("./html/*")

	a.s.GET("/ping", a.ping)

	a.s.GET("/", func(c *gin.Context) { a.page(c, "5TXCjnVNOE26Y8MtgIsNFh") })
	a.s.GET("/about", func(c *gin.Context) { a.page(c, "28O1On404p3JgNkCwgZ16q") })

	a.s.GET("/posts", a.posts)
	a.s.GET("/rss", a.rss)

	a.s.GET("/post/:slug/:id", a.post)

	a.s.NoRoute(a.notFound)

	return a
}

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

func (app app) Run() {
	app.s.Run()
}

func (app app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.s.ServeHTTP(w, r)
}
