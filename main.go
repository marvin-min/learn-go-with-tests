package main

import (
	"os"

	"github.com/marvin-min/learn-go-with-tests/di"
)

func main() {
	di.Greet(os.Stdout, "Eddile")
}
