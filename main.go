package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/marvin-min/learn-go-with-tests/di"
)

func main() {
	di.Greet(os.Stdout, "Eddile")
	str := "hhhhhhh饿了来了"
	sym := "h"
	for strings.HasPrefix(str, sym) {
		str = strings.TrimPrefix(str, sym)
	}
	fmt.Println(str)

}
