package app

import (
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

type cfgProvider interface {
	TemplateDir() string
	AboutID() string
	HomeID() string
}

type app struct {
	s   *gin.Engine
	c   contentProvider
	r   rssProvider
	cfg cfgProvider
}

func Default(cp contentProvider, rss rssProvider, cfg cfgProvider) app {
	a := app{gin.Default(), cp, rss, cfg}

	a.s.Static("/static", "./static")
	a.s.LoadHTMLGlob(cfg.TemplateDir())

	a.s.GET("/ping", a.ping)

	a.s.GET("/", func(c *gin.Context) { a.page(c, cfg.HomeID()) })
	a.s.GET("/about", func(c *gin.Context) { a.page(c, cfg.AboutID()) })

	a.s.GET("/posts", a.posts)
	a.s.GET("/rss", a.rss)

	a.s.GET("/post/:slug/:id", a.post)

	a.s.NoRoute(a.notFound)

	return a
}

func (app app) Run() {
	app.s.Run()
}

func (app app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.s.ServeHTTP(w, r)
}
