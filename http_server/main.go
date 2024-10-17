package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryStore()}
	log.Println("Starting player server on port :5000")
	log.Fatal(http.ListenAndServe(":5000", server))

}
