package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryStore())
	log.Println("Starting player server on port :5000")
	log.Fatal(http.ListenAndServe(":5000", server))

}
