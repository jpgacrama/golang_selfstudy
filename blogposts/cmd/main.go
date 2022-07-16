package main

import (
	"fmt"
	blogposts "github.com/jpgacrama/golang_selfstudy/blogposts"
	"log"
	"os"
)

func main() {
	fmt.Println("Running main.go. Getting Blogposts from Mocked FileSystem")
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
