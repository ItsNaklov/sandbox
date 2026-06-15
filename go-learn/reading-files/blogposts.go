// Package blogposts comment, as recommended
package blogposts

import "testing/fstest"

type Post struct{}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	return []Post{{}, {}}
}

// TODO write the refactor change the func
