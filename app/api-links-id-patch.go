package app

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/simpleserver"
)

func (a *App) ApiLinksIdPatch(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
