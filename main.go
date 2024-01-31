package main

import (
	"embed"
	"flag"
	"log"

	"github.com/andyinabox/linkydink/app"
)

//go:embed res/*
var res embed.FS

func main() {
	var host string
	var port string
	var dbfile string
	var mode string

	flag.StringVar(&host, "host", "127.0.0.1", "host to run the webserver on")
	flag.StringVar(&port, "port", "8000", "port to run the webserver on")
	flag.StringVar(&dbfile, "dbfile", "db/linkydink.db", "location on sqlite db")
	flag.StringVar(&mode, "mode", "debug", "run mode, use 'release' for production")
	flag.Parse()

	appInstance := app.New(&app.Config{
		Host:      host,
		Port:      port,
		Mode:      mode,
		DbFile:    dbfile,
		Resources: res,
	})

	log.Fatal(appInstance.Start())
}
