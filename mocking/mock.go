package mocking

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(writer, finalWord)
}
