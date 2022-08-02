package main

import (
	"bytes"
	"fmt"
	"golang_selfstudy/webApp"
	"os"
)

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	stdout := &bytes.Buffer{}
	game := &poker.Game{}
	poker.NewCLI(os.Stdin, stdout, game)
}
