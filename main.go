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
	var domain string
	var port string
	var dbfile string
	var mode string

	flag.StringVar(&domain, "domain", "linklink.click", "the domain the site is hosted on (linklink.click)")
	flag.StringVar(&port, "port", "", "port to run the webserver on")
	flag.StringVar(&dbfile, "dbfile", "db/linkydink.db", "location on sqlite db")
	flag.StringVar(&mode, "mode", "debug", "run mode, use 'release' for production")
	flag.Parse()

	lr := linkrepository.New(&linkrepository.Config{
		DbFile: dbfile,
	})
	ls := linkservice.New(lr)

	appInstance := app.New(&app.Config{
		Domain:    domain,
		Port:      port,
		Mode:      mode,
		Resources: res,
	}, ls)

	log.Fatal(appInstance.Start())
}
