package main

import (
	"fmt"
	"golang_selfstudy/webApp/src"
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
	blindAlerter := &poker.SpyBlindAlerter{}
	game := poker.NewTexasHoldem(blindAlerter, store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
	log.Fatalln("CLI version of main is not implemented yet")
}
