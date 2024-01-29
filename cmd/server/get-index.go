package main

import (
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := res.ReadFile("res/data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
