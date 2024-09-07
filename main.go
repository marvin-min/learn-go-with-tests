package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/marvin-min/learn-go-with-tests/clockface"
	"github.com/marvin-min/learn-go-with-tests/di"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

func others() {
	di.Greet(os.Stdout, "Eddile")
	str := "hhhhhhh饿了来了"
	sym := "h"
	for strings.HasPrefix(str, sym) {
		str = strings.TrimPrefix(str, sym)
	}
	fmt.Println(str)
}
