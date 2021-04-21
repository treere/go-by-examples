package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s*SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type CountdownOperationSpy struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Write message", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spysleeper := SpySleeper{}

		Countdown(&buffer, &spysleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Correct order", func(t *testing.T) {
		spy :=&CountdownOperationSpy{}

		Countdown(spy, spy)

		want := []string {
			write, sleep, write, sleep, write, sleep, write,
		}

		got := spy.Calls

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted calls %v got %v",  got, want)
		}
	})
}

func TestConfigurableSleep(t*testing.T) {
	sleepTime := 5*time.Second;

	spyTime :=&SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("shout hav slept %v but have slept %v", sleepTime, spyTime.durationSlept)
	}
	
}
