package main

import (
	"log"
	"net/http"
	"webApp/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
