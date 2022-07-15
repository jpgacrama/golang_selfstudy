package blogposts

import (
	"reflect"
	"regexp"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
					Description: Description 1
					Tags: tdd, go`
		secondBody = `Title: Post 2
					Description: Description 2
					Tags: rust, borrow-checker`
	)

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(replaceExtraSpaces(firstBody))},
		"hello-world2.md": {Data: []byte(replaceExtraSpaces(secondBody))},
	}
	posts, err := NewPostsFromFS(fs)
	assertDataIntegrityOfPosts(err, posts, t, fs)

	got := posts[0]
	want := Post{Title: "Post 1", Description: "Description 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	assertPost(t, posts[0], Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
	})
}

func assertDataIntegrityOfPosts(
	err error, posts []Post, t *testing.T, fs fstest.MapFS) {
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %+v posts, wanted %+v posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func replaceExtraSpaces(text string) string {
	space := regexp.MustCompile(`\t+`)
	textWithoutSpace := space.ReplaceAllString(text, "")
	return textWithoutSpace
}
