package main

import (
	"github.com/jpgacrama/golang_selfstudy/tree/webApp"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
