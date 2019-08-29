package data

/* page or post */
type Item interface {
	Title() string
	Desc() string
	Image() string
	Slug() string
	ID() string
}
