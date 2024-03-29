package blogrenderer

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"regexp"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

type PostRenderer struct {
	templ *template.Template
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
	bodySeparator        = "---"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			// todo: needs clarification, should we
			// totally fail if one file fails? or just ignore?
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func readMetaLine(scanner *bufio.Scanner, tagName string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), tagName)
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)
	return Post{
		Title:       readMetaLine(scanner, titleSeparator),
		Description: readMetaLine(scanner, descriptionSeparator),
		Tags:        strings.Split(readMetaLine(scanner, tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {

	if err := r.templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}

func ReplaceExtraSpaces(text string) string {
	space := regexp.MustCompile(`\t+`)
	textWithoutSpace := space.ReplaceAllString(text, "")
	return textWithoutSpace
}
