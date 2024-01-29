package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//go:embed res/*
var res embed.FS
var templates *template.Template

const domain = "127.0.0.1"
const port = "8000"

func init() {
	var err error
	templates, err = template.ParseFS(res, "res/*.tmpl")
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetIndex).Methods(http.MethodGet)
	http.Handle("/", router)

	fmt.Printf("Running server on %s:%s", domain, port)
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", domain, port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
