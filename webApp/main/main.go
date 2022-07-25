package main

import (
	"golang_selfstudy/webApp/playerstore"
	"golang_selfstudy/webApp/server"
	"log"
	"net/http"
)

func main() {
	server := &server.PlayerServer{Store: &playerstore.InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":8080", server))
}
