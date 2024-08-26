package ctx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello"
	svr := Server(&SpyStore{data})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	svr.ServeHTTP(resp, req)
	if resp.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, resp.Body.String(), data)
	}
}
