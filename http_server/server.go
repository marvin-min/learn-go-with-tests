package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(r, w)
	}

}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := parsePlayerName(r)
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func parsePlayerName(r *http.Request) string {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	return player
}

func (p *PlayerServer) showScore(r *http.Request, w http.ResponseWriter) {
	player := parsePlayerName(r)
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, score)
}
https://quii.gitbook.io/learn-go-with-tests/build-an-application/json