// Because you cannot use traditional templates on Now's Zeit 2.0. :(
package templates

import "html/template"

var (
	page = `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<title>https://evanjon.es/</title>
				<meta name="description" content="Personal website of Evan M Jones" />
				<meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
				<link href="https://fonts.googleapis.com/css?family=Abril+Fatface" rel="stylesheet" />
				<link href="/static/main.css" rel="stylesheet" />
				<script src="/static/main.js"></script>
			</head>
			<body>
				<aside>
					<div class="image-container">
						<img alt="{{ .sidebar.Fields.Alt }}" src="{{.sidebar.Fields.Image}}" />
					</div>
					<div class="content-container">{{ .sidebar.Fields.Raw }}</div>
				</aside>
				<main>
					<div class="image-container">
						<img alt="{{.item.Fields.Alt}}" src="{{.item.Fields.Image}}" />
					</div>
					<h1>{{ .item.Fields.Title }}</h1>
					<template>{{ .item.Fields.Desc }}</template>
				</main>
			</body>
		</html>
	`

	blog = `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<title>https://evanjon.es/</title>
				<meta name="description" content="Personal website of Evan M Jones" />
				<meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
				<link href="https://fonts.googleapis.com/css?family=Abril+Fatface" rel="stylesheet" />
				<link href="/static/main.css" rel="stylesheet" />
				<script src="/static/main.js"></script>
			</head>
			<body class="blog">
				<aside>
					<div class="image-container">
						<img alt="{{ .sidebar.Fields.Alt }}" src="{{.sidebar.Fields.Image}}" />
					</div>
					<div class="content-container">{{ .sidebar.Fields.Raw }}</div>
				</aside>
				<main>
					{{ range $key, $val := .items}}
					<a href="/post/{{ .Fields.Slug }}/{{ .Sys.ID }}">
						<img alt="{{.Fields.Alt}}" src="{{.Fields.Image}}" />
						<h1>{{ .Fields.Title }}</h1>
					</a>
					{{ end }}
				</main>
			</body>
		</html>
	`

	err = `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<title>https://evanjon.es/</title>
				<meta name="description" content="Personal website of Evan M Jones" />
				<meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
				<link href="https://fonts.googleapis.com/css?family=Abril+Fatface" rel="stylesheet" />
				<link href="/static/main.css" rel="stylesheet" />
				<script src="/static/main.js"></script>
			</head>
				<main>
					<center>
						<h1>Not found.</h1>
						<a href="/">Home</a>
					</center>
				</main>
			</body>
		</html>
	`
)

type Service struct {
	Root *template.Template
}

func Default() *Service {
	service := &Service{Root: template.New("")}
	service.Root.New("page.html").Parse(page)
	service.Root.New("blog.html").Parse(blog)
	service.Root.New("error.html").Parse(err)
	return service
}
