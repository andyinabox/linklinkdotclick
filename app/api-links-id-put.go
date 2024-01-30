package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/andyinabox/linkydink/pkg/simpleserver"
)

func (a *App) ApiLinksIdPut(ctx *simpleserver.Context) http.HandlerFunc {
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

		// look for link in database, if not send 404
		tx := a.db.First(&link, id)
		err = tx.Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// get link from response body
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&link)
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		// if the ids don't match, send bad request error
		if link.ID != uint(id) {
			ctx.WriteError(w, http.StatusBadRequest, errors.New("unmatched ids"))
			return
		}

		// refresh feed data if indicated
		// refreshing will automaticaly save
		if queryVars.Has("refresh") {
			err = a.RefreshLink(&link)
			if err != nil {
				ctx.WriteError(w, http.StatusInternalServerError, err)
				return
			}
			// otherwise save without refresh
		} else {
			tx = a.db.Save(&link)
			err = tx.Error
			if err != nil {
				ctx.WriteError(w, http.StatusInternalServerError, err)
				return
			}
		}

		// build response
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
