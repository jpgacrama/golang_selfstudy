package main

import (
	"golang_selfstudy/webApp/server"
	"log"
	"net/http"
)

func main() {
	server := &server.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
