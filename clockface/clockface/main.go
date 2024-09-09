package main

import (
	"os"
	"time"

	"github.com/marvin-min/learn-go-with-tests/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
