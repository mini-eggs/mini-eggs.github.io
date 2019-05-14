package handler

import (
	"net/http"
	"os"
	"evanjon.es/controller/content"
	"evanjon.es/service/contentful"
	"evanjon.es/service/templates"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	homeID := os.Getenv("HOME_ID")
	sidebarID := os.Getenv("SIDEBAR_ID")
	spaceID := os.Getenv("SPACE_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")

	Router := gin.Default()
	TemplateService := templates.Default()
	ContentfulService := contentful.Default(spaceID, accessToken)
	ContentController := content.Default(ContentfulService, sidebarID)

	Router.SetHTMLTemplate(TemplateService.Root)
	Router.GET("/", ContentController.SingleDefault(homeID))
	Router.GET("/post/:slug/:id", ContentController.Single)
	Router.GET("/posts", ContentController.List)

	return Router
}

func main() {
	r := router()
	r.Static("/static", "./static")
	r.Run()
}

// Handler - zeit
func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router().ServeHTTP(w, r)
}
