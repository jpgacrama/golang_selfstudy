package blogposts

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

type StubFailingFS struct {
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
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

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)
	return Post{
		Title:       readMetaLine(scanner, titleSeparator),
		Description: readMetaLine(scanner, descriptionSeparator),
		Tags:        strings.Split(readMetaLine(scanner, tagsSeparator), ", "),
	}, nil
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}
