package templatefuncs

import (
	"fmt"
	"html/template"
)

func CSS(css string) template.CSS {
	return template.CSS(css)
}

func HTML(html string) template.HTML {
	return template.HTML(html)
}

func Comment(text string) template.HTML {
	return template.HTML(fmt.Sprintf("<!-- %s -->", text))
}

func Funcs() template.FuncMap {
	return template.FuncMap{
		"css":     CSS,
		"html":    HTML,
		"comment": Comment,
	}
}

func NewWithFuncs(tmpl string) *template.Template {
	return template.New(tmpl).Funcs(Funcs())
}
