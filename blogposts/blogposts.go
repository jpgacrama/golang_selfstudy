package blogposts

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

type StubFailingFS struct {
}

const (
    titleSeparator       = "Title: "
    descriptionSeparator = "Description: "
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

func newPost(postFile io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postFile)

    readLine := func() string {
		scanner.Scan()
        return scanner.Text()
    }

    title := readLine()[len(titleSeparator):]
    description := readLine()[len(descriptionSeparator):]

    return Post{Title: title, Description: description}, nil
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}
