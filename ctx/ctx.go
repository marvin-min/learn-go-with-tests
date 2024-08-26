package ctx

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
}

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context#write-the-test-first