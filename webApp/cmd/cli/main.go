package main

import (
	"fmt"
	"golang_selfstudy/webApp"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	assertOpeningFileSuccessful(dbFileName, err)

	store, err := poker.NewFileSystemPlayerStore(db)
	assertCreatingFileSystemPlayerStoreSuccessful(err)

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}

func assertOpeningFileSuccessful(fileName string, err error) {
	if err != nil {
		log.Fatalf("problem opening %s %v", fileName, err)
	}
}

func assertCreatingFileSystemPlayerStoreSuccessful(err error) {
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
}
