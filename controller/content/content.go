package content

import (
	"net/http"

	"evanjon.es/service/contentful"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ContentfulService *contentful.Service
	sidebarID         string
}

func Default(ContentfulService *contentful.Service, sidebarID string) *Controller {
	return &Controller{
		ContentfulService: ContentfulService,
		sidebarID:         sidebarID,
	}
}

func (ctx *Controller) Single(c *gin.Context) {
	item, sidebar, err := ctx.withSidebar(c.Param("id"))

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "page.html", gin.H{"sidebar": sidebar, "item": item})

}

func (ctx *Controller) SingleDefault(id string) func(c *gin.Context) {
	return func(c *gin.Context) {
		item, sidebar, err := ctx.withSidebar(id)

		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{})
			return
		}

		c.HTML(http.StatusOK, "page.html", gin.H{"sidebar": sidebar, "item": item})
	}
}

func (ctx *Controller) List(c *gin.Context) {
	list := make(chan []contentful.Item)
	listErr := make(chan error)

	sidebar := make(chan contentful.Item)
	sidebarErr := make(chan error)

	go func() {
		items, err := ctx.ContentfulService.GetPosts()
		listErr <- err
		list <- items
	}()

	go func() {
		item, err := ctx.ContentfulService.GetPost(ctx.sidebarID)
		sidebarErr <- err
		sidebar <- item
	}()

	potentialErr1 := <-listErr
	potentialErr2 := <-sidebarErr

	if potentialErr1 != nil || potentialErr2 != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "blog.html", gin.H{"sidebar": <-sidebar, "items": <-list})
}

func (ctx *Controller) withSidebar(id string) (contentful.Item, contentful.Item, error) {
	item1 := make(chan contentful.Item)
	err1 := make(chan error)

	item2 := make(chan contentful.Item)
	err2 := make(chan error)

	go func() {
		item, err := ctx.ContentfulService.GetPost(id)
		err1 <- err
		item1 <- item
	}()

	go func() {
		item, err := ctx.ContentfulService.GetPost(ctx.sidebarID)
		err2 <- err
		item2 <- item
	}()
	
	potentialErr1 := <- err1
	if potentialErr1 != nil {
		return contentful.Item{}, contentful.Item{}, potentialErr1
	}

	potentialErr2 := <- err2
	if potentialErr2 != nil {
		return contentful.Item{}, contentful.Item{}, potentialErr2
	}

	return <-item1, <-item2, nil
}
