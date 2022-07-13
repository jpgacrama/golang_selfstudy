package blogposts

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello world2.md": {Data: []byte("hola")},
	}

	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %+v posts, wanted %+v posts", len(posts), len(fs))
	}

	got := posts[0]
	want := Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
