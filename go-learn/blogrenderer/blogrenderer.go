// Package blogrenderer asking me to write a comment find it annoying for some reason
package blogrenderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"

	"github.com/yuin/goldmark"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}
type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}

// Package embed provides access to files embedded in the running Go program.
//
//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := goldmark.Convert([]byte(p.Body), &body); err != nil {
		return err
	}

	view := struct {
		Title       string
		Description string
		Body        template.HTML
		Tags        []string
	}{
		Title:       p.Title,
		Description: p.Description,
		Body:        template.HTML(body.String()),
		Tags:        p.Tags,
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", view); err != nil {
		return err
	}

	return nil
}
