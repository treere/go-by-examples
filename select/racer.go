package selection

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b , 10 * time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timeout")
	}
}



func ping(url string) chan struct {} {
	ch := make(chan struct {})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
