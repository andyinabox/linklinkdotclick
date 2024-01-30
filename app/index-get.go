package app

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/simpleserver"
)

type RenderContext struct {
	Links []Link
}

func (a *App) IndexGet(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var links []Link
		tx := a.db.Order("last_clicked").Find(&links)
		err := tx.Error
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		renderContext := &RenderContext{
			Links: links,
		}

		// send response
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		// w.Write()
		ctx.Templates.ExecuteTemplate(w, "index.html.tmpl", renderContext)
	}
}
