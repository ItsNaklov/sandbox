// Package blogrenderer asking me to write a comment find it annoying for some reason
package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}

// TODOs: 1. Refactor testing file add t.Run
// 2 follow the tutorial carefully make sure to understand everything.
