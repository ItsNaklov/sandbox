// This is a package comment, as recommended
package blogposts

import "testing/fstest"

type Post struct{}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	return []Post{{}, {}}
}
