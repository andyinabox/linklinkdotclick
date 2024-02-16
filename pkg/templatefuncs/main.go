package templatefuncs

import "html/template"

func CSS(css string) template.CSS {
	return template.CSS(css)
}

func Funcs() template.FuncMap {
	return template.FuncMap{
		"css": CSS,
	}
}

func NewWithFuncs(tmpl string) *template.Template {
	return template.New(tmpl).Funcs(Funcs())
}
