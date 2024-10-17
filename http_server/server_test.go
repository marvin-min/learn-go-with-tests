package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("returns pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(response.Code, http.StatusOK, t)
		assetResponseBody(response.Body.String(), "20", t)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(response.Code, http.StatusOK, t)
		assetResponseBody(response.Body.String(), "10", t)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusNotFound
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{&store}
	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/player/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(response.Code, http.StatusAccepted, t)
	})
}

func assetResponseBody(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertStatus(got, want int, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}
func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
