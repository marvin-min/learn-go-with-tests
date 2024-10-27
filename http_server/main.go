package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryStore())
	log.Println("Starting player server on port http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", server))

}
