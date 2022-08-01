package main

import (
	"bytes"
	"fmt"
	"golang_selfstudy/webApp"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	var dummyStdOut = &bytes.Buffer{}
	poker.NewCLI(store, os.Stdin, dummyStdOut, poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()
}
