package read_files

import (
	"io/fs"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}

	return posts
}
