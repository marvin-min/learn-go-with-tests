package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(50 * time.Millisecond)
	return true
}
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
func TestWebsiteChecker(t *testing.T) {
	websites := []string{"https:/www.bing.com", "https://www.google.com", "waat://furhurterwe.geds"}
	want := map[string]bool{
		"https:/www.bing.com":     true,
		"https://www.google.com":  true,
		"waat://furhurterwe.geds": false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
