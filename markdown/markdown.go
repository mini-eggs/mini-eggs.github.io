package markdown

import (
	mder "gitlab.com/golang-commonmark/markdown"
)

type markdown struct {
	md *mder.Markdown
}

func Default() markdown {
	return markdown{mder.New(mder.XHTMLOutput(true))}
}

func (m markdown) Format(in string) (out string) {
	out = m.md.RenderToString([]byte(in))
	return
}
