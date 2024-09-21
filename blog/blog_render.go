package blog

import (
	"embed"
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"

	"io"
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}
func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.html", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	data := indexData{Posts: posts}
	return r.templ.ExecuteTemplate(w, "index.html", data)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

type indexData struct {
	Posts []Post
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
