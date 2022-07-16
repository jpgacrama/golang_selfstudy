package blogrenderer_test

import (
	"bytes"
	"github.com/jpgacrama/golang_selfstudy/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("Converts a single post into HTML", func(t *testing.T) {
		t.Helper()
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		t.Helper()
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>
				<p>This is a description</p>
				Tags: <ul><li>go</li><li>tdd</li></ul>`
		want = blogrenderer.ReplaceExtraSpaces(want)

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
