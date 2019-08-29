package contentful

import "html/template"

// item class, match data.Item interface

type item struct {
	theTitle, theDesc, theShortDesc, theImage, theImageAlt, theSlug, theID string
}

func (p item) Title() string {
	return p.theTitle
}

func (p item) Desc() template.HTML {
	return template.HTML(p.theDesc)
}

func (p item) ShortDesc() string {
	return p.theShortDesc
}

func (p item) Image() string {
	return p.theImage
}

func (p item) ImageAlt() string {
	return p.theImageAlt
}

func (p item) Slug() string {
	return p.theSlug
}

func (p item) ID() string {
	return p.theID
}
