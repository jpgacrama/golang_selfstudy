package blogposts

import (
	"errors"
	"io/fs"
	"testing/fstest"
)

type Post struct {
	Title string
}

type StubFailingFS struct {
}

func NewPostsFromFS(fileSystem fstest.MapFS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post

	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}
