package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/andyinabox/linkydink/pkg/simpleserver"
)

func (a *App) ApiLinksIdGet(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := ctx.Vars(r)
		queryVars, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			fmt.Printf("Error parsing query vars: %s", err.Error())
		}

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		var link Link
		tx := a.db.First(&link, id)
		err = tx.Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if queryVars.Has("refresh") {
			err = a.RefreshLink(&link)
			if err != nil {
				ctx.WriteError(w, http.StatusInternalServerError, err)
				return
			}
		}

		var responseData []byte
		responseData, err = json.Marshal(link)
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		// send response
		ctx.WriteJSON(w, responseData)
	}
}
