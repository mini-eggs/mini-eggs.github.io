package app

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"evanjon.es/data"
	"evanjon.es/rss"
	"github.com/stretchr/testify/assert"
)

// content provider

type acceptableContentProvider struct{}

func (c acceptableContentProvider) Get(id string) (ret data.Item, e error)    { return }
func (c acceptableContentProvider) List(id string) (ret []data.Item, e error) { return }

type brokenContentProvider struct{}

func (c brokenContentProvider) Get(id string) (ret data.Item, e error) {
	e = errors.New("failed to get post")
	return
}
func (c brokenContentProvider) List(id string) (ret []data.Item, e error) {
	e = errors.New("failed to get list")
	return
}

// config provider

type acceptableConfigProvider struct{}

func (cfg acceptableConfigProvider) TemplateDir() string { return "../html/*" }
func (cfg acceptableConfigProvider) HomeID() string      { return "" }
func (cfg acceptableConfigProvider) AboutID() string     { return "" }

// stub

func fakeEmptyRSS() rssProvider {
	return rss.Default("", "", "", "", "")
}

func fakeAcceptableApp() app {
	return Default(acceptableContentProvider{}, fakeEmptyRSS(), acceptableConfigProvider{})
}

func fakeBrokenApp() app {
	return Default(brokenContentProvider{}, fakeEmptyRSS(), acceptableConfigProvider{})
}

// tests

func TestPages(t *testing.T) {
	// home
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	app := fakeAcceptableApp()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "OK response is expected")

	// about
	req, _ = http.NewRequest("GET", "/about", nil)
	res = httptest.NewRecorder()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "OK response is expected")

	// posts
	req, _ = http.NewRequest("GET", "/posts", nil)
	res = httptest.NewRecorder()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "OK response is expected")

	// rss
	req, _ = http.NewRequest("GET", "/rss", nil)
	res = httptest.NewRecorder()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "OK response is expected")

	// 404 page
	req, _ = http.NewRequest("GET", "/does/not/exist", nil)
	res = httptest.NewRecorder()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code, "NotFound response is expected")

	// single post, exists
	req, _ = http.NewRequest("GET", "/post/some-slug/some-id", nil)
	res = httptest.NewRecorder()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "OK response is expected")

	// single post, does not exists
	app = fakeBrokenApp()
	req, _ = http.NewRequest("GET", "/post/some-slug/some-id", nil)
	res = httptest.NewRecorder()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code, "OK response is not expected")

	// posts, with failure
	req, _ = http.NewRequest("GET", "/posts", nil)
	res = httptest.NewRecorder()
	app.s.ServeHTTP(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code, "OK response is not expected")
}
