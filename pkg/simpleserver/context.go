package simpleserver

import (
	"embed"
	"html/template"
	"net/http"
)

type Context struct {
	Resources  embed.FS
	Templates  *template.Template
	Vars       func(*http.Request) map[string]string
	SetUrlVars func(*http.Request, map[string]string) *http.Request
	WriteError func(w http.ResponseWriter, code int, err error)
	WriteJSON  func(http.ResponseWriter, []byte)
}
