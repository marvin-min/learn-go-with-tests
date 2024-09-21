package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/marvin-min/learn-go-with-tests/di"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
func main() {
	html := mdToHTML([]byte("**hello world**"))
	fmt.Println(string(html))
}
func clock() {
	// t := time.Now()
	// clockface.SVGWriter(os.Stdout, t)
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

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/revisiting-arrays-and-slices-with-generics