package app

import (
	"net/http"

	"github.com/andyinabox/linkydink-sketch/pkg/simpleserver"
)

type ApiLinksPostData struct {
	Url string
}

func (a *App) ApiLinksPost(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
