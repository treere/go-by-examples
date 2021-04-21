package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c * ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (d* DefaultSleeper) Sleep() {
	time.Sleep(1* time.Second)
}

const finalWord = "Go!"
const countdownstart = 3

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := countdownstart; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}

	fmt.Fprint(buffer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1*time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
