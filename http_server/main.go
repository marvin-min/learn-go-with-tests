package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store := &FileSystemPlayerStore{db}
	server := NewPlayerServer(store)

	log.Println("Starting player server on port http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", server))

}
