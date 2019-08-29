package contentful

// item class, match data.Item interface

type item struct {
	theTitle, theDesc, theImage, theSlug, theID string
}

func (p item) Title() string {
	return p.theTitle
}

func (p item) Desc() string {
	return p.theDesc
}

func (p item) Image() string {
	return p.theImage
}

func (p item) Slug() string {
	return p.theSlug
}

func (p item) ID() string {
	return p.theID
}
