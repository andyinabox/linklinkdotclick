package app

import (
	"net/http"
	"strconv"

	"github.com/andyinabox/linkydink/pkg/simpleserver"
)

func (a *App) ApiLinksIdDelete(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := ctx.Vars(r)

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		tx := a.db.Delete(&Link{}, id)
		err = tx.Error
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
