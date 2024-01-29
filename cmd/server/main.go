package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

// go:embed res
var res embed.FS
var templates *template.Template

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetIndex).Methods(http.MethodGet)
	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
