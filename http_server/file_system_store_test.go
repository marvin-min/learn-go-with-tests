package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		store := FileSystemPlayerStore{database}
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(got, want, t)
		got = store.GetLeague()
		assertLeague(got, want, t)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		store := FileSystemPlayerStore{database}
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(got, want, t)
	})
}

func assertScoreEquals(got int, want int, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
https://quii.gitbook.io/learn-go-with-tests/build-an-application/io#write-the-test-first-2