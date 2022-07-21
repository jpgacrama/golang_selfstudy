package main

import (
	"golang_selfstudy/webApp/server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
