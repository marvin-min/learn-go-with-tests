package read_files

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("hei")},
		"hello_world2.md": {Data: []byte("hola")},
	}

	posts := NewPostFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files