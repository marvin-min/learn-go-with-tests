package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(writer, finalWord)
}
