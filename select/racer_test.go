package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("fast slow", func(t *testing.T) {
		slowServer := makeDealyerServer(20 * time.Millisecond)
		fastServer := makeDealyerServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if want != got {
			t.Errorf("Want %q got %q", want, got)
		}
	})

	t.Run("timeout", func( t*testing.T) {
		server := makeDealyerServer(25 * time.Millisecond)

		_, err := ConfigurableRacer(server.URL, server.URL, 15 * time.Millisecond)
		
		if err == nil {
			t.Error("Expected error")
		}
	})
}

func makeDealyerServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
