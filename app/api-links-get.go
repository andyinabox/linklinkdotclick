package app

import (
	"encoding/json"
	"net/http"

	"github.com/andyinabox/linkydink/pkg/simpleserver"
)

func (a *App) ApiLinksGet(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var links []Link
		tx := a.db.Find(&links)
		err := tx.Error
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		var responseData []byte
		responseData, err = json.Marshal(links)
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		ctx.WriteJSON(w, responseData)
	}
}
