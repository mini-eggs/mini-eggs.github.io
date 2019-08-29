package handler

import (
	"net/http"
	"os"

	"evanjon.es/app"
	"evanjon.es/config"
	"evanjon.es/contentful"
	"evanjon.es/rss"
)

type I interface {
	Run()
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func build() I {
	c := contentful.Default(
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

func main() {
	build().Run()
}

func H(w http.ResponseWriter, r *http.Request) {
	build().ServeHTTP(w, r)
}
