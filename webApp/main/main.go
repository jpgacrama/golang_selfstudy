package main

import (
	"github.com/jpgacrama/golang_selfstudy/webApp"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
