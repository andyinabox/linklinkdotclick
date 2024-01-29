package main

import (
	"fmt"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s", "World")
}
