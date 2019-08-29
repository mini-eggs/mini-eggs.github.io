package data

import "html/template"

/* page or post */
type Item interface {
	Title() string
	Desc() template.HTML
	Image() string
	ImageAlt() string
	Slug() string
	ID() string
}
