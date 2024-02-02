package main

import (
	"embed"
	"flag"
	"log"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/app/linkservice"
)

//go:embed res/*
var res embed.FS

func main() {
	var host string
	var port string
	var dbfile string
	var mode string
	// var tz string

	flag.StringVar(&host, "host", "127.0.0.1", "host to run the webserver on")
	flag.StringVar(&port, "port", "8000", "port to run the webserver on")
	flag.StringVar(&dbfile, "dbfile", "db/linkydink.db", "location on sqlite db")
	flag.StringVar(&mode, "mode", "debug", "run mode, use 'release' for production")
	// flag.StringVar(&tz, "tz", "Europe/Madrid", "time zone, default 'Europe/Madrid")
	flag.Parse()

	lr := linkrepository.New(&linkrepository.Config{
		DbFile: dbfile,
	})
	ls := linkservice.New(lr)

	appInstance := app.New(&app.Config{
		Host: host,
		Port: port,
		Mode: mode,
		// TimeZone:  tz,
		Resources: res,
	}, ls)

	log.Fatal(appInstance.Start())
}
