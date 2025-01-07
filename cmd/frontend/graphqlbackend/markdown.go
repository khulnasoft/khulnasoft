package graphqlbackend

import "github.com/khulnasoft/khulnasoft/internal/markdown"

type Markdown string

func (m Markdown) Text() string {
	return string(m)
}

func (m Markdown) HTML() (string, error) {
	return markdown.Render(string(m))
}
