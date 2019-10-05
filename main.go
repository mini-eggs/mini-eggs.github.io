package handler

import (
	"log"
	"net/http"
	"os"

	"evanjon.es/app"
	"evanjon.es/config"
	"evanjon.es/contentful"
	"evanjon.es/markdown"
	"evanjon.es/rss"
)

type server interface {
	Run() error
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func build() server {
	c := contentful.Default(
		markdown.Default(),
		os.Getenv("SPACE_ID"),
		os.Getenv("TOKEN"),
	)

	r := rss.Default(
		"evanjon.es",
		"My personal corner of the big WWW.",
		"https://evanjon.es/",
		"Evan M Jones",
		"me@evanjon.es",
	)

	cfg := config.Default(
		os.Getenv("HOME_ID"),
		os.Getenv("ABOUT_ID"),
	)

	return app.Default(c, r, cfg)
}

// This func is never used in production.
// Error checking ultimately doesn't matter here.
func main() {
	if e := build().Run(); e != nil {
		log.Fatal(e)
	}
}

// For Zeit's Now 2.0.
func H(w http.ResponseWriter, r *http.Request) {
	build().ServeHTTP(w, r)
}
