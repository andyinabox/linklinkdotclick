package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/app/linkservice"
	"github.com/joho/godotenv"
)

//go:embed res/*
var res embed.FS

func main() {
	var host string
	var port string
	var dbfile string
	var mode string
	// var tz string

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .dnv file")
	}

	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "db/linkydink.db"
	}

	flag.StringVar(&host, "host", os.Getenv("HOST"), "host to run the webserver on")
	flag.StringVar(&port, "port", os.Getenv("PORT"), "port to run the webserver on")
	flag.StringVar(&dbfile, "dbfile", dbFile, "location on sqlite db")
	flag.StringVar(&mode, "mode", os.Getenv("GIN_MODE"), "run mode, use 'release' for production")
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
