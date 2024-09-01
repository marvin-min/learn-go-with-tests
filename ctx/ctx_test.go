package ctx

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	// if !s.cancelled {
	// 	s.t.Error("store was not told to cancel")
	// }
}
func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	// if s.cancelled {
	// 	s.t.Error("store was told to cancel")
	// }
}
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case r := <-data:
		return r, nil

	}
}
func (s *SpyStore) Cancel() {
	// s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancelCtx)

		resp := &SpyResponseWriter{}

		svr.ServeHTTP(resp, req)

		if resp.written {
			t.Error("a response should not have been written")
		}
	})
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		if resp.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, resp.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})
}
