package main

import (
	"embed"
	"flag"
	"log"
	"path"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/app/userservice"
	"github.com/gin-contrib/sessions/cookie"
)

//go:embed res/*
var res embed.FS

func main() {
	var domain string
	var port string
	var dbfile string
	var mode string
	var defaultemail string

	flag.StringVar(&domain, "domain", "linklink.click", "the domain the site is hosted on (linklink.click)")
	flag.StringVar(&port, "port", "8080", "port to run the webserver on")
	flag.StringVar(&dbfile, "dbfile", "db/linkydink.db", "location on sqlite db")
	flag.StringVar(&mode, "mode", "debug", "run mode, use 'release' for production")
	flag.StringVar(&defaultemail, "defaultemail", "linkydink@linkydink.tld", "an email for the default user that appears when not logged in")
	flag.Parse()

	// maybe not the most secure?
	store := cookie.NewStore([]byte(domain + port + dbfile + mode + defaultemail))

	userRepository, err := userrepository.New(&userrepository.Config{
		DbFile: dbfile,
	})
	if err != nil {
		panic(err)
	}

	userService := userservice.New(&userservice.Config{
		UserDbPath:       path.Join(path.Dir(dbfile), "usr"),
		DefaultUserEmail: defaultemail,
	}, userRepository)

	appInstance := app.New(&app.Config{
		Domain:    domain,
		Port:      port,
		Mode:      mode,
		Resources: res,
	}, userService, store)

	log.Fatal(appInstance.Start())
}
