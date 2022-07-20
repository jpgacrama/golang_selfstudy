package main

import (
	"log"
	"net/http"
	"webApp"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
