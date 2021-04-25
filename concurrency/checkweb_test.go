package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "http://pippo.calippo"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.it",
		"http://pippo.calippo",
		"http://youtube.com",
	}

	want := map[string]bool{
		"http://google.it":     true,
		"http://pippo.calippo": false,
		"http://youtube.com":   true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
